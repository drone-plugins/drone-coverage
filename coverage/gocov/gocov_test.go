package gocov

import (
	"reflect"
	"testing"

	"github.com/drone-plugins/drone-coverage/coverage"
	"golang.org/x/tools/cover"
)

func TestParse(t *testing.T) {
	got, err := New().Read(sampleFile)
	if err != nil {
		t.Errorf("Expected Go coverage profile parsed successfully, got error %s", err)
	}

	if !reflect.DeepEqual(got, sampleProfiles) {
		t.Errorf("Expected Go coverage profile matches the test fixture")
	}
}

func TestSniff(t *testing.T) {
	ok, _ := coverage.FromBytes([]byte("foo:"))
	if ok {
		t.Errorf("Expect sniffer does not find match")
	}

	ok, r := coverage.FromBytes(sampleFile)
	if !ok {
		t.Errorf("Expect sniffer to find a match")
	}
	if _, ok := r.(*reader); !ok {
		t.Errorf("Expect sniffer to return a reader")
	}
}

var sampleProfiles = []*cover.Profile{
	{
		FileName: "github.com/drone-plugins/drone-coverage/coverage/lcov/lcov.go",
		Mode:     "set",
		Blocks: []cover.ProfileBlock{
			{Count: 1, StartLine: 14, StartCol: 50, EndLine: 17, EndCol: 2, NumStmt: 2},
			{Count: 1, StartLine: 20, StartCol: 52, EndLine: 26, EndCol: 15, NumStmt: 5},
			{Count: 1, StartLine: 26, StartCol: 15, EndLine: 29, EndCol: 10, NumStmt: 2},
			{Count: 1, StartLine: 30, StartCol: 3, EndLine: 34, EndCol: 40, NumStmt: 4},
			{Count: 1, StartLine: 35, StartCol: 3, EndLine: 50, EndCol: 50, NumStmt: 11},
			{Count: 1, StartLine: 51, StartCol: 3, EndLine: 52, EndCol: 12, NumStmt: 1},
			{Count: 1, StartLine: 56, StartCol: 2, EndLine: 56, EndCol: 22, NumStmt: 1},
		},
	},
}

var sampleFile = []byte(`mode: set
github.com/drone-plugins/drone-coverage/coverage/lcov/lcov.go:14.50,17.2 2 1
github.com/drone-plugins/drone-coverage/coverage/lcov/lcov.go:20.52,26.15 5 1
github.com/drone-plugins/drone-coverage/coverage/lcov/lcov.go:56.2,56.22 1 1
github.com/drone-plugins/drone-coverage/coverage/lcov/lcov.go:26.15,29.10 2 1
github.com/drone-plugins/drone-coverage/coverage/lcov/lcov.go:30.3,34.40 4 1
github.com/drone-plugins/drone-coverage/coverage/lcov/lcov.go:35.3,50.50 11 1
github.com/drone-plugins/drone-coverage/coverage/lcov/lcov.go:51.3,52.12 1 1`)
