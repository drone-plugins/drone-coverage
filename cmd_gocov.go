package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/param108/drone-coverage/coverage/gocov"
	"github.com/mattn/go-zglob"
	"github.com/urfave/cli"
	"golang.org/x/tools/cover"
)

// GocovCmd is the exported command for converting Go coverage files.
var GocovCmd = cli.Command{
	Name:  "gocov",
	Usage: "parse gocov files",
	Flags: []cli.Flag{},
	Action: func(c *cli.Context) error {
		return parseGocov(c)
	},
}

func parseGocov(c *cli.Context) error {

	pattern := c.Args().First()
	if pattern == "" {
		pattern = "**/*.out"
	}

	matches, err := zglob.Glob(pattern)
	if err != nil {
		return err
	}

	parser := gocov.New()
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
