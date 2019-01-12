package main

import (
	"os"
	"testing"

	"github.com/drone-plugins/drone-coverage/client"
	"github.com/sirupsen/logrus"
)

func createReport(names ...string) *client.Report {
	report := &client.Report{}
	var files []*client.File

	for _, name := range names {
		file := &client.File{
			FileName: name,
		}

		files = append(files, file)
	}

	report.Files = files
	return report
}

func TestPathModification(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stderr)

	// test modified
	report := createReport("a/b.go")
	findFileReferences(report, "/a/")

	fileName := report.Files[0].FileName
	if fileName != "b.go" {
		t.Errorf("Expected filename to be b.go not %s", fileName)
	}

	// test removed
	report = createReport("a/b.go", "b/c.go")
	findFileReferences(report, "/a/")

	if len(report.Files) != 1 {
		t.Errorf("Expected a file to be removed")
	} else {
		fileName = report.Files[0].FileName
		if fileName != "b.go" {
			t.Errorf("Expected filename to be b.go not %s", fileName)
		}
	}

	// test file present
	f, err := os.Create(".a.go")

	if err != nil {
		t.Errorf("Could not create file for test")
	}
	defer os.Remove(f.Name())

	report = createReport(".a.go")
	findFileReferences(report, "/a/")

	if len(report.Files) != 1 {
		t.Errorf("Expected file to be present")
	} else {
		fileName = report.Files[0].FileName
		if fileName != ".a.go" {
			t.Errorf("Expected filename to be .a.go not %s", fileName)
		}
	}
}
