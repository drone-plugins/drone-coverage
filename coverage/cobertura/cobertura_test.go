package cobertura

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/drone-plugins/drone-coverage/coverage"

	"golang.org/x/tools/cover"
)

func TestParseXMLToStruct(t *testing.T) {
	c, err := parse(sampleFile)

	if err != nil {
		t.Fatalf("Expected Go cobertura parser to parse XML file successfully, got error %s", err)
	}

	if !reflect.DeepEqual(c, sampleStruct) {
		t.Errorf("Expected Go cobertura parsed struct to match the test fixture")
	}
}

func TestParse(t *testing.T) {
	got, err := New().Read(sampleFile)

	if err != nil {
		t.Fatalf("Expected Go coverage profile parsed successfully, got error %s", err)
	}

	if !reflect.DeepEqual(got, sampleProfiles) {
		t.Errorf("Expected Go coverage profile matches the test fixture")
	}
}

func TestParseFromFile(t *testing.T) {
	got, err := New().ReadFile(sampleFile2)

	if err != nil {
		t.Fatalf("Expected Go coverage profile parsed successfully, got error %s", err)
	}

	if !reflect.DeepEqual(got, sampleProfiles2) {
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
		t.Fatalf("Expect sniffer to find a match")
	}
	if _, ok := r.(*reader); !ok {
		t.Errorf("Expect sniffer to return a Cobertura reader")
	}
}

var sampleFile = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE coverage SYSTEM "http://cobertura.sourceforge.net/xml/coverage-03.dtd">
<coverage line-rate="0" branch-rate="0" version="" timestamp="1476758263048">
	<packages>
		<package name="github.com/drone-plugins/drone-coverage/coverage/gocov" line-rate="0" branch-rate="0" complexity="0">
			<classes>
				<class name="-" filename="/go/src/github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go" line-rate="0" branch-rate="0" complexity="0">
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
				<class name="reader" filename="/go/src/github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go" line-rate="0" branch-rate="0" complexity="0">
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

var sampleProfiles = []*cover.Profile{
	{
		FileName: "/go/src/github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go",
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

var sampleStruct = cobertura{
	XMLName: xml.Name{Local: "coverage"},
	Classes: []class{
		{
			Filename: "/go/src/github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go",
			Lines: []line{
				{Number: 14, Hits: 1},
				{Number: 21, Hits: 1},
			},
		},
		{
			Filename: "/go/src/github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go",
			Lines: []line{
				{Number: 25, Hits: 1},
				{Number: 26, Hits: 1},
				{Number: 30, Hits: 1},
				{Number: 34, Hits: 1},
				{Number: 35, Hits: 1},
				{Number: 36, Hits: 0},
				{Number: 38, Hits: 1},
				{Number: 39, Hits: 1},
				{Number: 40, Hits: 0},
				{Number: 42, Hits: 1},
			},
		},
	},
}

var sampleFile2 = "coverage-sample.xml"

var sampleProfiles2 = []*cover.Profile{
	{
		FileName: "Main.java",
		Mode:     "set",
		Blocks: []cover.ProfileBlock{
			{Count: 3, StartLine: 10, EndLine: 10, NumStmt: 1},
			{Count: 3, StartLine: 16, EndLine: 16, NumStmt: 1},
			{Count: 3, StartLine: 17, EndLine: 17, NumStmt: 1},
			{Count: 3, StartLine: 18, EndLine: 18, NumStmt: 1},
			{Count: 3, StartLine: 19, EndLine: 19, NumStmt: 1},
			{Count: 3, StartLine: 23, EndLine: 23, NumStmt: 1},
			{Count: 3, StartLine: 25, EndLine: 25, NumStmt: 1},
			{Count: 3, StartLine: 26, EndLine: 26, NumStmt: 1},
			{Count: 3, StartLine: 28, EndLine: 28, NumStmt: 1},
			{Count: 3, StartLine: 29, EndLine: 29, NumStmt: 1},
			{Count: 3, StartLine: 30, EndLine: 30, NumStmt: 1},
		},
	},
	{
		FileName: "search/BinarySearch.java",
		Mode:     "set",
		Blocks: []cover.ProfileBlock{
			{Count: 3, StartLine: 12, EndLine: 12, NumStmt: 1},
			{Count: 3, StartLine: 16, EndLine: 16, NumStmt: 1},
			{Count: 12, StartLine: 18, EndLine: 18, NumStmt: 1},
			{Count: 9, StartLine: 20, EndLine: 20, NumStmt: 1},
			{Count: 9, StartLine: 21, EndLine: 21, NumStmt: 1},
			{Count: 9, StartLine: 23, EndLine: 23, NumStmt: 1},
			{Count: 0, StartLine: 24, EndLine: 24, NumStmt: 1},
			{Count: 9, StartLine: 25, EndLine: 25, NumStmt: 1},
			{Count: 6, StartLine: 26, EndLine: 26, NumStmt: 1},
			{Count: 3, StartLine: 28, EndLine: 28, NumStmt: 1},
			{Count: 9, StartLine: 29, EndLine: 29, NumStmt: 1},
			{Count: 3, StartLine: 31, EndLine: 31, NumStmt: 1},
		},
	},
	{
		FileName: "search/ISortedArraySearch.java",
		Mode:     "set",
	},
	{
		FileName: "search/LinearSearch.java",
		Mode:     "set",
		Blocks: []cover.ProfileBlock{
			{Count: 3, StartLine: 9, EndLine: 9, NumStmt: 1},
			{Count: 9, StartLine: 13, EndLine: 13, NumStmt: 1},
			{Count: 9, StartLine: 15, EndLine: 15, NumStmt: 1},
			{Count: 3, StartLine: 16, EndLine: 16, NumStmt: 1},
			{Count: 6, StartLine: 17, EndLine: 17, NumStmt: 1},
			{Count: 0, StartLine: 19, EndLine: 19, NumStmt: 1},
			{Count: 0, StartLine: 24, EndLine: 24, NumStmt: 1},
		},
	},
}
