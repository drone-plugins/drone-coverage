package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/drone-plugins/drone-coverage/client"
	"github.com/drone-plugins/drone-coverage/coverage"
	"github.com/joho/godotenv"
	"github.com/mattn/go-zglob"
	"github.com/urfave/cli"
	"golang.org/x/tools/cover"
)

// PublishCmd is the exported command for publishing coverage files.
var PublishCmd = cli.Command{
	Name:  "publish",
	Usage: "publish coverage report",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "repo.fullname",
			Usage:  "repository full name",
			EnvVar: "DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "git commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "commit.author.name",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.author.avatar",
			Usage:  "git author avatar",
			EnvVar: "DRONE_COMMIT_AUTHOR_AVATAR",
		},
		cli.StringFlag{
			Name:   "pattern",
			Usage:  "coverage file pattern",
			Value:  "**/*.*",
			EnvVar: "PLUGIN_PATTERN",
		},
		cli.StringFlag{
			Name:   "server",
			Usage:  "coverage server",
			Value:  "**/*.*",
			EnvVar: "PLUGIN_SERVER",
		},
		cli.Float64Flag{
			Name:   "threshold",
			Usage:  "coverage threshold",
			EnvVar: "PLUGIN_THRESHOLD",
		},
		cli.BoolFlag{
			Name:   "increase",
			Usage:  "coverage must increase",
			EnvVar: "PLUGIN_MUST_INCREASE",
		},
		cli.StringFlag{
			Name:   "cert",
			Usage:  "coverage cert",
			EnvVar: "COVERAGE_CERT",
		},
		cli.StringFlag{
			Name:   "token",
			Usage:  "github token",
			EnvVar: "PLUGIN_TOKEN,GITHUB_TOKEN",
		},
		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
		cli.StringFlag{
			Name:   "trim-prefix",
			Usage:  "trim prefix from coverage files",
			EnvVar: "PLUGIN_TRIM_PREFIX",
		},
	},
	Action: func(c *cli.Context) error {
		return publish(c)
	},
}

func publish(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	logrus.Debugf("finding coverage files that match %s", c.String("pattern"))

	matches, err := zglob.Glob(c.String("pattern"))
	if err != nil {
		return err
	}

	var profiles []*cover.Profile
	for _, match := range matches {
		ok, reader := coverage.FromFile(match)
		if !ok {
			continue
		}

		logrus.Debugf("found coverage file %s", match)

		parsed, rerr := reader.ReadFile(match)
		if rerr != nil {
			return err
		}

		for _, p := range parsed {
			profiles = addProfile(profiles, p)
		}
	}

	// create the coverage payload that gets sent to the
	// coverage reporting server.
	report := profileToReport(profiles)

	build := client.Build{
		Number:    c.Int("build.number"),
		Event:     c.String("build.event"),
		Commit:    c.String("commit.sha"),
		Branch:    c.String("commit.branch"),
		Ref:       c.String("commit.ref"),
		Message:   c.String("commit.message"),
		Author:    c.String("commit.author.name"),
		Avatar:    c.String("commit.author.avatar"),
		Link:      c.String("build.link"),
		Timestamp: time.Now().UTC().Unix(),
	}

	// get the base directory
	base := c.String("trim.prefix")

	if len(base) == 0 {
		base, err = os.Getwd()

		if err != nil {
			return err
		}

		logrus.Debug("Using current working directory")
	}

	logrus.Debugf("Base directory is %s", base)

	findFileReferences(report, base)

	var (
		repo   = c.String("repo.fullname")
		server = c.String("server")
		secret = c.String("token")
		cert   = c.String("cert")
	)

	cli := newClient(server, cert, "")
	token, err := cli.Token(secret)
	if err != nil {
		return err
	}
	cli = newClient(server, cert, token.Access)

	// check and see if the repository exists. if not, activate
	if _, err := cli.Repo(repo); err != nil {
		if _, err := cli.Activate(repo); err != nil {
			return err
		}
	}

	resp, err := cli.Submit(repo, &build, report)
	if err != nil {
		return err
	}

	switch {
	case resp.Changed > 0:
		fmt.Printf("Code coverage increased %.1f%% to %.1f%%\n", resp.Changed, resp.Coverage)
	case resp.Changed < 0:
		fmt.Printf("Code coverage dropped %.1f%% to %.1f%%\n", resp.Changed, resp.Coverage)
	default:
		fmt.Printf("Code coverage unchanged, %.1f%%\n", resp.Coverage)
	}

	if c.Float64("threshold") < resp.Coverage && c.Float64("threshold") != 0 {
		return fmt.Errorf("Failing build. Coverage threshold may not fall below %.1f%%\n", c.Float64("threshold"))
	}
	if resp.Changed < 0 && c.Bool("increase") {
		return fmt.Errorf("Failing build. Coverage may not decrease")
	}

	return nil
}

func findFileReferences(report *client.Report, base string) {
	var files []*client.File

	// normalize the file path based on the working directory
	// also ignore any files outside of the directory
	for _, file := range report.Files {
		fileName := file.FileName
		var prefix string

		if path.IsAbs(fileName) {
			if !strings.HasPrefix(fileName, base) {
				logrus.Warningf("File referenced in coverage not found at %s", fileName)
				continue
			}

			prefix = base
		} else if _, err := os.Stat(fileName); err == nil {
			logrus.Debugf("File found at relative path %s", fileName)

			prefix = ""
		} else {
			var err error
			prefix, err = coverage.PathPrefix(fileName, base)

			if err != nil {
				// See if file is on disk
				logrus.Warningf("File referenced in coverage not found at %s", fileName)
				continue
			}

			logrus.Debugf("Found common path at %s", prefix)
		}

		// Add the file to the report
		file.FileName = strings.TrimPrefix(fileName, prefix)
		files = append(files, file)
	}

	report.Files = files
}

// profileToReport is a helper function that converts the merged coverage
// report to the Report JSON format expected by the coverage server.
func profileToReport(profiles []*cover.Profile) *client.Report {
	report := client.Report{}
	report.Files = make([]*client.File, len(profiles), len(profiles))

	for i, profile := range profiles {
		file := client.File{
			Mode:     profile.Mode,
			FileName: profile.FileName,
		}

		file.Blocks = make([]*client.Block, len(profile.Blocks), len(profile.Blocks))
		for ii, block := range profile.Blocks {
			file.Blocks[ii] = &client.Block{
				StartLine: block.StartLine,
				StartCol:  block.StartCol,
				EndLine:   block.EndLine,
				EndCol:    block.EndCol,
				NumStmt:   block.NumStmt,
				Count:     block.Count,
			}
		}

		covered, total, percent := percentCovered(profile)
		file.Lines = total
		file.Covered = covered
		file.Coverage = percent

		report.Files[i] = &file
		report.Lines += file.Lines
		report.Covered += file.Covered
	}
	if report.Lines != 0 {
		report.Coverage = float64(report.Covered) / float64(report.Lines) * float64(100)
	}
	return &report
}

// percentCovered is a helper fucntion that calculate the percent
// coverage for coverage profile.
func percentCovered(p *cover.Profile) (int64, int64, float64) {
	var total, covered int64
	for _, b := range p.Blocks {
		total += int64(b.NumStmt)
		if b.Count > 0 {
			covered += int64(b.NumStmt)
		}
	}
	var percent float64
	if total != 0 {
		percent = float64(covered) / float64(total) * float64(100)
	}
	return covered, total, percent
}

// newClient returns a new coverage server client.
func newClient(server, cert, token string) client.Client {
	pool, err := x509.SystemCertPool()
	if err != nil {
		pool = x509.NewCertPool()
	}
	conf := &tls.Config{RootCAs: pool}
	pem, _ := ioutil.ReadFile(cert)
	if len(pem) != 0 {
		pool.AppendCertsFromPEM(pem)
	}
	if len(token) == 0 {
		return client.NewClientTLS(server, conf)
	}
	return client.NewClientTokenTLS(server, token, conf)
}
