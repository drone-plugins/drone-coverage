package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// trace is a simple debugging function to print the next executed command.
func trace(cmd *exec.Cmd) {
	fmt.Printf("+ %s\n", strings.Join(cmd.Args, " "))
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

		file.Covered, file.Lines, file.Coverage = percentCovered(profile)

		report.Files[i] = &file
		report.Lines += file.Lines
		report.Covered += file.Covered
	}

	if report.Lines != 0 {
		report.Coverage = float64(report.Covered) / float64(report.Lines) * float64(100)
	}

	return &report
}

// percentCovered is a helper fucntion that calculate the percent coverage for
// coverage profile.
func percentCovered(p *cover.Profile) (int64, int64, float64) {
	var (
		percent float64
		total   int64
		covered int64
	)

	for _, b := range p.Blocks {
		total += int64(b.NumStmt)

		if b.Count > 0 {
			covered += int64(b.NumStmt)
		}
	}

	if total != 0 {
		percent = float64(covered) / float64(total) * float64(100)
	}

	return covered, total, percent
}

// resolveHost is a helper function that returns the default coverage server url.
func resolveServer(raw string) string {
	result, err := url.Parse(raw)

	if err != nil {
		return "https://aircover.co"
	}

	host, _, err := net.SplitHostPort(result.Host)

	if err != nil {
		host = result.Host
	}

	items, err := net.LookupTXT(host)

	if err != nil {
		return "https://aircover.co"
	}

	for _, txt := range items {
		if strings.HasPrefix(txt, "coverage=") {
			return txt[9:]
		}
	}

	return "https://aircover.co"
}
