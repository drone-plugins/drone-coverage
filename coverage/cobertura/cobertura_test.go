package cobertura

import (
	"reflect"
	"testing"

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

// mode: set
// github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go:13.13,15.2 1 1
// github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go:20.28,22.2 1 1
// github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go:24.61,27.2 2 1
// github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go:29.66,31.2 1 1
// github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go:33.68,35.16 2 1
// github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go:38.2,39.46 2 1
// github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go:42.2,42.32 1 1
// github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go:35.16,37.3 1 0
// github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go:39.46,41.3 1 0
var sampleProfiles = []*cover.Profile{
	{
		FileName: "github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go",
		Mode:     "set",
		Blocks: []cover.ProfileBlock{
			{Count: 1, StartLine: 13, StartCol: 13, EndLine: 15, EndCol: 2, NumStmt: 1},
			{Count: 1, StartLine: 20, StartCol: 28, EndLine: 22, EndCol: 2, NumStmt: 1},
			{Count: 1, StartLine: 24, StartCol: 61, EndLine: 27, EndCol: 2, NumStmt: 2},
			{Count: 1, StartLine: 29, StartCol: 66, EndLine: 31, EndCol: 2, NumStmt: 1},
			{Count: 1, StartLine: 33, StartCol: 68, EndLine: 35, EndCol: 16, NumStmt: 2},
			{Count: 1, StartLine: 38, StartCol: 3, EndLine: 39, EndCol: 46, NumStmt: 2},
			{Count: 1, StartLine: 42, StartCol: 2, EndLine: 42, EndCol: 32, NumStmt: 1},
		},
	},
}

var sampleFile = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE coverage SYSTEM "http://cobertura.sourceforge.net/xml/coverage-03.dtd">
<coverage line-rate="0" branch-rate="0" version="" timestamp="1476758263048">
	<packages>
		<package name="github.com/drone-plugins/drone-coverage/coverage/gocov" line-rate="0" branch-rate="0" complexity="0">
			<classes>
				<class name="-" filename="/home/fbcbarbosa/Development/go/src/github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go" line-rate="0" branch-rate="0" complexity="0">
					<methods>
						<method name="init" signature="" line-rate="0" branch-rate="0">
							<lines>
								<line number="14" hits="1"></line>
							</lines>
						</method>
						<method name="New" signature="" line-rate="0" branch-rate="0">
							<lines>
								<line number="21" hits="1"></line>
							</lines>
						</method>
					</methods>
					<lines>
						<line number="14" hits="1"></line>
						<line number="21" hits="1"></line>
					</lines>
				</class>
				<class name="reader" filename="/home/fbcbarbosa/Development/go/src/github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go" line-rate="0" branch-rate="0" complexity="0">
					<methods>
						<method name="Read" signature="" line-rate="0" branch-rate="0">
							<lines>
								<line number="25" hits="1"></line>
								<line number="26" hits="1"></line>
							</lines>
						</method>
						<method name="ReadFile" signature="" line-rate="0" branch-rate="0">
							<lines>
								<line number="30" hits="1"></line>
							</lines>
						</method>
						<method name="ReadFrom" signature="" line-rate="0" branch-rate="0">
							<lines>
								<line number="34" hits="1"></line>
								<line number="35" hits="1"></line>
								<line number="36" hits="0"></line>
								<line number="38" hits="1"></line>
								<line number="39" hits="1"></line>
								<line number="40" hits="0"></line>
								<line number="42" hits="1"></line>
							</lines>
						</method>
					</methods>
					<lines>
						<line number="25" hits="1"></line>
						<line number="26" hits="1"></line>
						<line number="30" hits="1"></line>
						<line number="34" hits="1"></line>
						<line number="35" hits="1"></line>
						<line number="36" hits="0"></line>
						<line number="38" hits="1"></line>
						<line number="39" hits="1"></line>
						<line number="40" hits="0"></line>
						<line number="42" hits="1"></line>
					</lines>
				</class>
			</classes>
		</package>
	</packages>
</coverage>`)
