package cobertura

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/drone-plugins/drone-coverage/coverage"

	"golang.org/x/tools/cover"
)

func TestParseXMLToStruct(t *testing.T) {
	c, err := parse(sampleGolangToCobertura)

	if err != nil {
		t.Fatalf("Expected Go cobertura parser to parse XML file successfully, got error %s", err)
	}

	if !reflect.DeepEqual(c, sampleGolangToCoberturaStructs) {
		t.Errorf("Expected Go cobertura parsed struct to match the test fixture")
	}
}

func TestParseGolangToCobertura(t *testing.T) {
	got, err := New().Read(sampleGolangToCobertura)

	if err != nil {
		t.Fatalf("Expected Go coverage profile parsed successfully, got error %s", err)
	}

	if !reflect.DeepEqual(got, sampleGolangToCoberturaProfile) {
		t.Errorf("Expected Go coverage profile matches the test fixture")
	}
}

func TestParsePythonCoverage(t *testing.T) {
	got, err := New().Read(sampleJunitCoverage)

	if err != nil {
		t.Fatalf("Expected Go coverage profile parsed successfully, got error %s", err)
	}

	if !reflect.DeepEqual(got, sampleJunitCoverageProfile) {
		t.Errorf("Expected Go coverage profile matches the test fixture")
	}
}

func TestSniff(t *testing.T) {
	ok, _ := coverage.FromBytes([]byte("foo:"))
	if ok {
		t.Errorf("Expect sniffer does not find match")
	}

	ok, r := coverage.FromBytes(sampleGolangToCobertura)
	if !ok {
		t.Fatalf("Expect sniffer to find a match")
	}
	if _, ok := r.(*reader); !ok {
		t.Errorf("Expect sniffer to return a Cobertura reader")
	}

	ok, r = coverage.FromBytes(sampleJunitCoverage)
	if !ok {
		t.Fatalf("Expect sniffer to find a match")
	}
	if _, ok = r.(*reader); !ok {
		t.Errorf("Expect sniffer to return a Cobertura reader")
	}
}

var sampleGolangToCobertura = []byte(`<?xml version="1.0" encoding="UTF-8"?>
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

var sampleGolangToCoberturaProfile = []*cover.Profile{
	{
		FileName: "/go/src/github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go",
		Mode:     "set",
		Blocks: []cover.ProfileBlock{
			{Count: 1, StartLine: 14, StartCol: 0, EndLine: 14, EndCol: 0, NumStmt: 1},
			{Count: 1, StartLine: 21, StartCol: 0, EndLine: 21, EndCol: 0, NumStmt: 1},
			{Count: 1, StartLine: 25, StartCol: 0, EndLine: 25, EndCol: 0, NumStmt: 1},
			{Count: 1, StartLine: 26, StartCol: 0, EndLine: 26, EndCol: 0, NumStmt: 1},
			{Count: 1, StartLine: 30, StartCol: 0, EndLine: 30, EndCol: 0, NumStmt: 1},
			{Count: 1, StartLine: 34, StartCol: 0, EndLine: 34, EndCol: 0, NumStmt: 1},
			{Count: 1, StartLine: 35, StartCol: 0, EndLine: 35, EndCol: 0, NumStmt: 1},
			{Count: 0, StartLine: 36, StartCol: 0, EndLine: 36, EndCol: 0, NumStmt: 1},
			{Count: 1, StartLine: 38, StartCol: 0, EndLine: 38, EndCol: 0, NumStmt: 1},
			{Count: 1, StartLine: 39, StartCol: 0, EndLine: 39, EndCol: 0, NumStmt: 1},
			{Count: 0, StartLine: 40, StartCol: 0, EndLine: 40, EndCol: 0, NumStmt: 1},
			{Count: 1, StartLine: 42, StartCol: 0, EndLine: 42, EndCol: 0, NumStmt: 1},
		},
	},
}

var sampleGolangToCoberturaStructs = cobertura{
	XMLName: xml.Name{Local: "coverage"},
	Classes: []class{
		{
			FileName: "/go/src/github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go",
			Lines: []line{
				{Number: 14, Hits: 1},
				{Number: 21, Hits: 1},
			},
		},
		{
			FileName: "/go/src/github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go",
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

var sampleJunitCoverage = []byte(`<?xml version="1.0"?>
<!--DOCTYPE coverage SYSTEM "http://cobertura.sourceforge.net/xml/coverage-03.dtd"-->

<coverage line-rate="0.9" branch-rate="0.75" version="1.9" timestamp="1187350905008">
	<sources>
		<source>C:/local/mvn-coverage-example/src/main/java</source>
		<source>--source</source>
	</sources>
	<packages>
		<package name="" line-rate="1.0" branch-rate="1.0" complexity="1.0">
			<classes>
				<class name="Main" filename="Main.java" line-rate="1.0" branch-rate="1.0" complexity="1.0">
					<methods>
						<method name="&lt;init&gt;" signature="()V" line-rate="1.0" branch-rate="1.0">
							<lines>
								<line number="10" hits="3" branch="false"/>
							</lines>
						</method>
						<method name="doSearch" signature="()V" line-rate="1.0" branch-rate="1.0">
							<lines>
								<line number="23" hits="3" branch="false"/>
								<line number="25" hits="3" branch="false"/>
								<line number="26" hits="3" branch="false"/>
								<line number="28" hits="3" branch="false"/>
								<line number="29" hits="3" branch="false"/>
								<line number="30" hits="3" branch="false"/>
							</lines>
						</method>
						<method name="main" signature="([Ljava/lang/String;)V" line-rate="1.0" branch-rate="1.0">
							<lines>
								<line number="16" hits="3" branch="false"/>
								<line number="17" hits="3" branch="false"/>
								<line number="18" hits="3" branch="false"/>
								<line number="19" hits="3" branch="false"/>
							</lines>
						</method>
					</methods>
					<lines>
						<line number="10" hits="3" branch="false"/>
						<line number="16" hits="3" branch="false"/>
						<line number="17" hits="3" branch="false"/>
						<line number="18" hits="3" branch="false"/>
						<line number="19" hits="3" branch="false"/>
						<line number="23" hits="3" branch="false"/>
						<line number="25" hits="3" branch="false"/>
						<line number="26" hits="3" branch="false"/>
						<line number="28" hits="3" branch="false"/>
						<line number="29" hits="3" branch="false"/>
						<line number="30" hits="3" branch="false"/>
					</lines>
				</class>
			</classes>
		</package>
		<package name="search" line-rate="0.8421052631578947" branch-rate="0.75" complexity="3.25">
			<classes>
				<class name="search.BinarySearch" filename="search/BinarySearch.java" line-rate="0.9166666666666666" branch-rate="0.8333333333333334" complexity="3.0">
					<methods>
						<method name="&lt;init&gt;" signature="()V" line-rate="1.0" branch-rate="1.0">
							<lines>
								<line number="12" hits="3" branch="false"/>
							</lines>
						</method>
						<method name="find" signature="([II)I" line-rate="0.9090909090909091" branch-rate="0.8333333333333334">
							<lines>
								<line number="16" hits="3" branch="false"/>
								<line number="18" hits="12" branch="true" condition-coverage="100% (2/2)">
									<conditions>
										<condition number="0" type="jump" coverage="100%"/>
									</conditions>
								</line>
								<line number="20" hits="9" branch="false"/>
								<line number="21" hits="9" branch="false"/>
								<line number="23" hits="9" branch="true" condition-coverage="50% (1/2)">
									<conditions>
										<condition number="0" type="jump" coverage="50%"/>
									</conditions>
								</line>
								<line number="24" hits="0" branch="false"/>
								<line number="25" hits="9" branch="true" condition-coverage="100% (2/2)">
									<conditions>
										<condition number="0" type="jump" coverage="100%"/>
									</conditions>
								</line>
								<line number="26" hits="6" branch="false"/>
								<line number="28" hits="3" branch="false"/>
								<line number="29" hits="9" branch="false"/>
								<line number="31" hits="3" branch="false"/>
							</lines>
						</method>
					</methods>
					<lines>
						<line number="12" hits="3" branch="false"/>
						<line number="16" hits="3" branch="false"/>
						<line number="18" hits="12" branch="true" condition-coverage="100% (2/2)">
							<conditions>
								<condition number="0" type="jump" coverage="100%"/>
							</conditions>
						</line>
						<line number="20" hits="9" branch="false"/>
						<line number="21" hits="9" branch="false"/>
						<line number="23" hits="9" branch="true" condition-coverage="50% (1/2)">
							<conditions>
								<condition number="0" type="jump" coverage="50%"/>
							</conditions>
						</line>
						<line number="24" hits="0" branch="false"/>
						<line number="25" hits="9" branch="true" condition-coverage="100% (2/2)">
							<conditions>
								<condition number="0" type="jump" coverage="100%"/>
							</conditions>
						</line>
						<line number="26" hits="6" branch="false"/>
						<line number="28" hits="3" branch="false"/>
						<line number="29" hits="9" branch="false"/>
						<line number="31" hits="3" branch="false"/>
					</lines>
				</class>
				<class name="search.ISortedArraySearch" filename="search/ISortedArraySearch.java" line-rate="1.0" branch-rate="1.0" complexity="1.0">
					<methods>
					</methods>
					<lines>
					</lines>
				</class>
				<class name="search.LinearSearch" filename="search/LinearSearch.java" line-rate="0.7142857142857143" branch-rate="0.6666666666666666" complexity="6.0">
					<methods>
						<method name="&lt;init&gt;" signature="()V" line-rate="1.0" branch-rate="1.0">
							<lines>
								<line number="9" hits="3" branch="false"/>
							</lines>
						</method>
						<method name="find" signature="([II)I" line-rate="0.6666666666666666" branch-rate="0.6666666666666666">
							<lines>
								<line number="13" hits="9" branch="true" condition-coverage="50% (1/2)">
									<conditions>
										<condition number="0" type="jump" coverage="50%"/>
									</conditions>
								</line>
								<line number="15" hits="9" branch="true" condition-coverage="100% (2/2)">
									<conditions>
										<condition number="0" type="jump" coverage="100%"/>
									</conditions>
								</line>
								<line number="16" hits="3" branch="false"/>
								<line number="17" hits="6" branch="true" condition-coverage="50% (1/2)">
									<conditions>
										<condition number="0" type="jump" coverage="50%"/>
									</conditions>
								</line>
								<line number="19" hits="0" branch="false"/>
								<line number="24" hits="0" branch="false"/>
							</lines>
						</method>
					</methods>
					<lines>
						<line number="9" hits="3" branch="false"/>
						<line number="13" hits="9" branch="true" condition-coverage="50% (1/2)">
							<conditions>
								<condition number="0" type="jump" coverage="50%"/>
							</conditions>
						</line>
						<line number="15" hits="9" branch="true" condition-coverage="100% (2/2)">
							<conditions>
								<condition number="0" type="jump" coverage="100%"/>
							</conditions>
						</line>
						<line number="16" hits="3" branch="false"/>
						<line number="17" hits="6" branch="true" condition-coverage="50% (1/2)">
							<conditions>
								<condition number="0" type="jump" coverage="50%"/>
							</conditions>
						</line>
						<line number="19" hits="0" branch="false"/>
						<line number="24" hits="0" branch="false"/>
					</lines>
				</class>
			</classes>
		</package>
	</packages>
</coverage>`)

var sampleJunitCoverageProfile = []*cover.Profile{
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
