package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/drone-plugins/drone-coverage/client"
	"github.com/drone-plugins/drone-coverage/coverage"
	"github.com/mattn/go-zglob"
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
		cli.StringFlag{
			Name:   "cert",
			Usage:  "coverage cert",
			EnvVar: "COVERAGE_CERT",
		},
		cli.StringFlag{
			Name:   "token",
			Usage:  "coverage token",
			EnvVar: "COVERAGE_TOKEN",
		},
	},
	Action: func(c *cli.Context) {
		err := publish(c)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func publish(c *cli.Context) error {

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
		parsed, err := reader.ReadFile(match)
		if err != nil {
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

	// this code attempts we use the relative path to the
	// project instead of an absoluate path. We should probably
	// just exclude anything not in the repository workspace ...
	// for _, file := range report.Files {
	// 	// convert from absolute to relative path
	// 	file.FileName = strings.TrimPrefix(
	// 		file.FileName,
	// 		w.Path,
	// 	)
	// 	// convert from gopath to relative path
	// 	file.FileName = strings.TrimPrefix(
	// 		file.FileName,
	// 		strings.TrimPrefix(w.Path, "/drone/src/"),
	// 	)
	// 	// remove report prefix
	// 	file.FileName = strings.TrimPrefix(file.FileName, "/")
	// }

	// Use the GitHub token in the Netrc file to authenticate
	// to the coverage server. For security purposes, we only
	// do this for the official coverage service.
	// if p.Config.Token == "" && p.Config.Server == "" {
	// 	v.Token = w.Netrc.Login
	// }
	// if v.Server == "" {
	// 	// v.Server = resolveServer(s.Link)
	// }

	// Handle provided custom CA certs
	caCertPool := x509.NewCertPool()
	tlsConfig := &tls.Config{RootCAs: caCertPool}
	cli := client.NewClient(c.String("server"))
	if c.String("cert") != "" {
		caCertPool.AppendCertsFromPEM([]byte(c.String("cert")))
		tlsConfig = &tls.Config{RootCAs: caCertPool}
		cli = client.NewClientTLS(c.String("server"), tlsConfig)
	}

	token, err := cli.Token(c.String("token"))
	if err != nil {
		return err
	}
	if c.String("cert") != "" {
		cli = client.NewClientTokenTLS(c.String("server"), token.Access, tlsConfig)
	} else {
		cli = client.NewClientToken(c.String("server"), token.Access)
	}

	// check and see if the repository exists. if not, activate
	if _, err := cli.Repo(c.String("repo.fullname")); err != nil {
		if _, err := cli.Activate(c.String("repo.fullname")); err != nil {
			return err
		}
	}

	resp, err := cli.Submit(c.String("repo.fullname"), &build, report)
	if err != nil {
		return err
	}

	if resp != nil {
		// nothing. this is just here to avoid unused variable compiler error for now
	}

	// switch {
	// case resp.Changed > 0:
	// 	return fmt.Errorf("Code coverage increased %.1f%% to %.1f%%\n", resp.Changed, resp.Coverage)
	// case resp.Changed < 0:
	// 	return fmt.Errorf("Code coverage dropped %.1f%% to %.1f%%\n", resp.Changed, resp.Coverage)
	// default:
	// 	return fmt.Errorf("Code coverage unchanged, %.1f%%\n", resp.Coverage)
	// }

	// if p.Config.Threshold < resp.Coverage && p.Config.Threshold != 0 {
	// 	return fmt.Errorf("Failing build. Coverage threshold may not fall below %.1f%%\n", p.Config.Threshold)
	// }
	// if resp.Changed < 0 && p.Config.MustIncrease {
	// 	return fmt.Errorf("Failing build. Coverage may not decrease")
	// }

	return nil
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
