package main

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/tools/cover"

	"github.com/codegangsta/cli"
	"github.com/drone-plugins/drone-coverage/coverage/lcov"
	"github.com/mattn/go-zglob"
)

// LcovCmd is the exported command for converting LCOV files.
var LcovCmd = cli.Command{
	Name:  "lcov",
	Usage: "parse lcov files",
	Flags: []cli.Flag{},
	Action: func(c *cli.Context) {
		err := parseLcov(c)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func parseLcov(c *cli.Context) error {

	pattern := c.Args().First()
	if pattern == "" {
		pattern = "**/lcov.info"
	}

	matches, err := zglob.Glob(pattern)
	if err != nil {
		return err
	}

	parser := lcov.New()
	var profiles []*cover.Profile
	for _, match := range matches {
		parsed, err := parser.ReadFile(match)
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

	out, err := json.MarshalIndent(report, " ", " ")
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "%s", out)
	return nil
}
