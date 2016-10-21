package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/drone-plugins/drone-coverage/coverage/cobertura"
	"github.com/mattn/go-zglob"
	"github.com/urfave/cli"
	"golang.org/x/tools/cover"
)

// CoberturaCmd is the exported command for converting Cobertura files.
var CoberturaCmd = cli.Command{
	Name:  "cobertura",
	Usage: "parse cobertura files",
	Flags: []cli.Flag{},
	Action: func(c *cli.Context) error {
		return parseCobertura(c)
	},
}

func parseCobertura(c *cli.Context) error {

	pattern := c.Args().First()
	if pattern == "" {
		pattern = "**/coverage.xml"
	}

	matches, err := zglob.Glob(pattern)
	if err != nil {
		return err
	}

	parser := cobertura.New()
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
