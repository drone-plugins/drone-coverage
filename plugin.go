package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/drone-plugins/drone-coverage/client"
	"github.com/drone-plugins/drone-coverage/coverage"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
	"golang.org/x/tools/cover"

	_ "github.com/drone-plugins/drone-coverage/coverage/gocov"
	_ "github.com/drone-plugins/drone-coverage/coverage/lcov"
)

type Plugin struct {
	Repo   Repo
	Build  Build
	Netrc  Netrc
	Config Config
}

func (p Plugin) Exec() error {
	var (
		merged  []*cover.Profile
		include *regexp.Regexp
	)

	if p.Config.Include != "" {
		include, _ = regexp.CompilePOSIX(p.Config.Include)

		if include == nil {
			return fmt.Errorf("Error compiling regular expression %s\n", p.Config.Include)
		}
	}

	filepath.Walk(w.Path, func(path string, info os.FileInfo, err error) error {
		// attempt to match the coverage file by path name using either
		// the default regular expression or the custom.
		if include != nil && !include.MatchString(path) {
			return nil
		} else if !coverage.IsMatch(path) {
			return nil
		}

		fmt.Printf("Parsing coverage file %s\n", path)
		ok, reader := coverage.FromFile(path)

		if !ok {
			fmt.Printf("Failure to determine coverage format %s\n", path)
			return nil
		}

		profiles, err := reader.ReadFile(path)

		if err != nil {
			return err
		}

		for _, p := range profiles {
			merged = addProfile(merged, p)
		}

		return nil
	})

	// create the coverage payload that gets sent to the
	// coverage reporting server.
	report := profileToReport(merged)

	// TODO(must) Replace that???
	// DRONE_BUILD_NUMBER
	// DRONE_BUILD_EVENT
	// DRONE_COMMIT_SHA
	// DRONE_COMMIT_BRANCH
	// DRONE_COMMIT_REF
	// ?
	// DRONE_COMMIT_MESSAGE
	// DRONE_COMMIT_AUTHOR
	// DRONE_COMMIT_AUTHOR_AVATAR
	// DRONE_BUILD_LINK

	// build := client.Build{
	//   Number:    b.Number,
	//   Event:     b.Event,
	//   Commit:    b.Commit,
	//   Branch:    b.Branch,
	//   Ref:       b.Ref,
	//   Refspec:   b.Refspec,
	//   Message:   b.Message,
	//   Author:    b.Author,
	//   Avatar:    b.Avatar,
	//   Link:      b.Link,
	//   Timestamp: time.Now().UTC().Unix(),
	// }

	// this code attempts we use the relative path to the
	// project instead of an absoluate path. We should probably
	// just exclude anything not in the repository workspace ...
	for _, file := range report.Files {
		// convert from absolute to relative path
		file.FileName = strings.TrimPrefix(
			file.FileName,
			"w.Path", // TODO(must): What is the replacement for w.Path?
		)

		// convert from gopath to relative path
		file.FileName = strings.TrimPrefix(
			file.FileName,
			strings.TrimPrefix("w.Path", "/drone/src/"), // TODO(must): What is the replacement for w.Path?
		)

		// remove report prefix
		file.FileName = strings.TrimPrefix(file.FileName, "/")
	}

	// Use the GitHub token in the Netrc file to authenticate
	// to the coverage server. For security purposes, we only
	// do this for the official coverage service.
	if p.Config.Token == "" && p.Config.Server == "" {
		p.Config.Token = p.Netrc.Login
	}

	if p.Config.Server == "" {
		p.Config.Server = resolveServer(p.Build.Link)
	}

	caCertPool := x509.NewCertPool()
	tlsConfig := &tls.Config{RootCAs: caCertPool}

	cli := client.NewClient(p.Config.Server)

	if p.Config.CACert != "" {
		caCertPool.AppendCertsFromPEM([]byte(p.Config.CACert))
		tlsConfig = &tls.Config{RootCAs: caCertPool}

		cli = client.NewClientTLS(p.Config.Server, tlsConfig)
	}

	token, err := cli.Token(p.Config.Token)

	if err != nil {
		return fmt.Errorf("Cannot authenticate. %s\n", err)
	}

	if p.Config.CACert != "" {
		cli = client.NewClientTokenTLS(p.Config.Server, token.Access, tlsConfig)
	} else {
		cli = client.NewClientToken(p.Config.Server, token.Access)
	}

	if _, err := cli.Repo(r.Repo.FullName); err != nil {
		if _, err := cli.Activate(r.Repo.FullName); err != nil {
			return fmt.Errorf("Cannot activate repository. %s\n", err)
		}
	}

	resp, err := cli.Submit(r.Repo.FullName, &build, report)

	if err != nil {
		return fmt.Errorf("Cannot submit coverage. %s\n", err)
	}

	switch {
	case resp.Changed > 0:
		fmt.Printf("Code coverage increased %.1f%% to %.1f%%\n", resp.Changed, resp.Coverage)
	case resp.Changed < 0:
		fmt.Printf("Code coverage dropped %.1f%% to %.1f%%\n", resp.Changed, resp.Coverage)
	default:
		fmt.Printf("Code coverage unchanged, %.1f%%\n", resp.Coverage)
	}

	if p.Config.Threshold < resp.Coverage && p.Config.Threshold != 0 {
		return fmt.Errorf("Failing build. Coverage threshold may not fall below %.1f%%\n", p.Config.Threshold)
	}

	if resp.Changed < 0 && p.Config.MustIncrease {
		return fmt.Errorf("Failing build. Coverage may not decrease")
	}

	return nil
}
