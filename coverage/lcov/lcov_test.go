package lcov

import (
	"reflect"
	"testing"

	"github.com/drone-plugins/drone-coverage/coverage"
	"golang.org/x/tools/cover"
)

func TestRead(t *testing.T) {
	got, err := New().Read(sampleFile)
	if err != nil {
		t.Errorf("Expected LCOV parsed successfully, got error %s", err)
	}

	if !reflect.DeepEqual(got, sampleProfiles) {
		t.Errorf("Expected LCOV parsed file equals test fixture")
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
		t.Errorf("Expect sniffer to return an LCOV reader")
	}
}

func TestRelativePathsReplacement(t *testing.T) {
	got, err := New().Read(sampleFileWithRelativePaths)
	if err != nil {
		t.Errorf("Expected LCOV parsed successfully, got error %s", err)
	}

	if !reflect.DeepEqual(got, sampleProfilesWithAbsolutePaths) {
		t.Errorf("Expected LCOV parsed file equals test fixture")
	}
}

var sampleProfilesWithAbsolutePaths = []*cover.Profile{
	{
		FileName: "/drone/src/github.com/donny-dont/dogma-codegen/lib/src/codegen/function_generator.dart",
		Mode:     "set",
		Blocks: []cover.ProfileBlock{
			{NumStmt: 1, StartLine: 30, EndLine: 30, Count: 1},
		},
	},
}

var sampleProfiles = []*cover.Profile{
	{
		FileName: "/drone/src/github.com/donny-dont/dogma-codegen/lib/src/codegen/function_generator.dart",
		Mode:     "set",
		Blocks: []cover.ProfileBlock{
			{NumStmt: 1, StartLine: 30, EndLine: 30, Count: 1},
			{NumStmt: 1, StartLine: 31, EndLine: 31, Count: 1},
			{NumStmt: 1, StartLine: 32, EndLine: 32, Count: 1},
			{NumStmt: 1, StartLine: 34, EndLine: 34, Count: 1},
			{NumStmt: 1, StartLine: 42, EndLine: 42, Count: 0},
			{NumStmt: 1, StartLine: 45, EndLine: 45, Count: 1},
			{NumStmt: 1, StartLine: 48, EndLine: 48, Count: 1},
			{NumStmt: 1, StartLine: 52, EndLine: 52, Count: 1},
			{NumStmt: 1, StartLine: 54, EndLine: 54, Count: 1},
			{NumStmt: 1, StartLine: 58, EndLine: 58, Count: 1},
			{NumStmt: 1, StartLine: 62, EndLine: 62, Count: 1},
		},
	},
	{
		FileName: "/drone/src/github.com/donny-dont/dogma-codegen/lib/src/codegen/annotated_metadata_generator.dart",
		Mode:     "set",
		Blocks: []cover.ProfileBlock{
			{NumStmt: 1, StartLine: 33, EndLine: 33, Count: 1},
			{NumStmt: 1, StartLine: 36, EndLine: 36, Count: 1},
			{NumStmt: 1, StartLine: 37, EndLine: 37, Count: 1},
			{NumStmt: 1, StartLine: 38, EndLine: 38, Count: 1},
		},
	},
}

var sampleFileWithRelativePaths = []byte(`
SF:./drone/src/github.com/donny-dont/dogma-codegen/lib/src/codegen/function_generator.dart
DA:30,84
end_of_record
`)

var sampleFile = []byte(`
SF:/drone/src/github.com/donny-dont/dogma-codegen/lib/src/codegen/function_generator.dart
DA:30,84
DA:31,28
DA:32,56
DA:34,56
DA:42,0
DA:45,28
DA:48,28
DA:52,12
DA:54,16
DA:58,28
DA:62,16
end_of_record
SF:/drone/src/github.com/donny-dont/dogma-codegen/lib/src/codegen/annotated_metadata_generator.dart
DA:33,230
DA:36,479
DA:37,236
DA:38,51
end_of_record
`)
