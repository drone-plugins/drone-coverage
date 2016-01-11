package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/drone-plugins/drone-coverage/client"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
	"golang.org/x/tools/cover"
)

type params struct {
	Server       string  `json:"server"`
	Token        string  `json:"token"`
	Threshold    float64 `json:"threshold"`
	MustIncrease bool    `json:"must_increase"`
}

func main() {
	var (
		w = drone.Workspace{}
		b = drone.Build{}
		r = drone.Repo{}
		v = params{}
	)
	plugin.Param("workspace", &w)
	plugin.Param("build", &b)
	plugin.Param("repo", &r)
	plugin.Param("vargs", &v)
	plugin.MustParse()

	var merged []*cover.Profile

	// merge all coverage reports into a single report
	var walker = func(path string, info os.FileInfo, err error) error {
		match, _ := filepath.Match("*.out", filepath.Base(path))
		if !match {
			return nil
		}
		fmt.Printf("Parsing coverage file %s\n", path)
		profiles, err := cover.ParseProfiles(path)
		if err != nil {
			return err
		}
		for _, p := range profiles {
			merged = addProfile(merged, p)
		}
		return nil
	}
	filepath.Walk(w.Path, walker)

	// create the coverage payload that gets sent to the
	// coverage reporting server.
	report := profileToReport(merged)
	build := client.Build{
		Number:  b.Number,
		Event:   b.Event,
		Commit:  b.Commit,
		Branch:  b.Branch,
		Ref:     b.Ref,
		Refspec: b.Refspec,
		Message: b.Message,
		Author:  b.Author,
		Avatar:  b.Avatar,
		Link:    b.Link,
	}

	// this code attempts we use the relative path to the
	// project instead of an absoluate path. We should probably
	// just exclude anything not in the repository workspace ...
	for _, file := range report.Files {
		file.FileName = strings.TrimPrefix(
			file.FileName,
			strings.TrimPrefix(w.Path, "/drone/src/"),
		)
		file.FileName = strings.TrimPrefix(file.FileName, "/")
	}

	// Use the GitHub token in the Netrc file to authenticate
	// to the coverage server. For security purposes, we only
	// do this for the official coverage service.
	if v.Token == "" && v.Server == "" {
		v.Token = w.Netrc.Login
	}
	if v.Server == "" {
		v.Server = "https://coverage.drone.io"
	}

	cli := client.NewClient(v.Server)
	token, err := cli.Token(v.Token)
	if err != nil {
		fmt.Printf("Cannot authenticate. %s\n", err)
		os.Exit(1)
	}
	cli = client.NewClientToken(v.Server, token.Access)

	// check and see if the repository exists. if not, activate
	if _, err := cli.Repo(r.FullName); err != nil {
		if _, err := cli.Activate(r.FullName); err != nil {
			fmt.Printf("Cannot activate repository. %s\n", err)
			os.Exit(1)
		}
	}

	resp, err := cli.Submit(r.FullName, &build, report)
	if err != nil {
		fmt.Printf("Cannot submit coverage. %s\n", err)
		os.Exit(1)
	}

	switch {
	case resp.Changed > 0:
		fmt.Printf("Code coverage increased %.1f%% to %.1f%%\n", resp.Changed, resp.Coverage)
	case resp.Changed < 0:
		fmt.Printf("Code coverage dropped %.1f%% to %.1f%%\n", resp.Changed, resp.Coverage)
	default:
		fmt.Printf("Code coverage unchanged, %.1f%%\n", resp.Coverage)
	}

	if v.Threshold < resp.Coverage {
		fmt.Printf("Failing build. Coverage threshold may not fall below %.1f%%\n", v.Threshold)
		os.Exit(1)
	}
	if resp.Changed < 0 {
		fmt.Println("Failing build. Coverage may not decrease")
		os.Exit(1)
	}
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

// percentCovered is a help fucntion that calculate the percent
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
