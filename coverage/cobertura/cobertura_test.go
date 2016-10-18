package cobertura

import (
	"fmt"
	"testing"

	"golang.org/x/tools/cover"
)

func TestXML(t *testing.T) {
	r := new(reader)
	x, err := r.parseXML(sampleFile)

	fmt.Println(x)

	if err != nil {
		t.Fatalf("Expected XML Cobertura file parsed successfully, got error %s", err)
	}

	if x.Packages[0].Classes[0].Filename != "/home/fbcbarbosa/Development/go/src/github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go" {
		t.Errorf("Wrong name, got %s", x.Packages[0].Classes[0].Name)
	}
}

// func TestParse(t *testing.T) {
// 	got, err := New().Read(sampleFile)
// 	if err != nil {
// 		t.Errorf("Expected Go coverage profile parsed successfully, got error %s", err)
// 	}

// 	if !reflect.DeepEqual(got, sampleProfiles) {
// 		t.Errorf("Expected Go coverage profile matches the test fixture")
// 	}
// }

var sampleProfiles = []*cover.Profile{
	{
		FileName: "/home/fbcbarbosa/Development/go/src/github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go",
		Mode:     "set",
		Blocks: []cover.ProfileBlock{
			{Count: 1, StartLine: 14, EndLine: 14, NumStmt: 1},
			{Count: 1, StartLine: 21, EndLine: 21, NumStmt: 1},
			{Count: 1, StartLine: 25, EndLine: 25, NumStmt: 1},
			{Count: 1, StartLine: 26, EndLine: 26, NumStmt: 1},
			{Count: 1, StartLine: 30, EndLine: 30, NumStmt: 1},
			{Count: 1, StartLine: 34, EndLine: 34, NumStmt: 1},
			{Count: 1, StartLine: 35, EndLine: 35, NumStmt: 1},
			{Count: 0, StartLine: 36, EndLine: 36, NumStmt: 1},
			{Count: 1, StartLine: 38, EndLine: 38, NumStmt: 1},
			{Count: 1, StartLine: 39, EndLine: 39, NumStmt: 1},
			{Count: 0, StartLine: 40, EndLine: 40, NumStmt: 1},
			{Count: 1, StartLine: 42, EndLine: 42, NumStmt: 1},
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
