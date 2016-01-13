package coverage

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"regexp"

	"golang.org/x/tools/cover"
)

// the algorithm uses at most sniffLen bytes to make its decision.
const sniffLen = 128

// the list of registered sniffers.
var sniffers = map[Reader][]byte{}

// regular expression that can be used to match common filepath patterns
// as a first check to determine if something is a coverage report.
var includes = regexp.MustCompile(".+\\.out|.+clover.xml|.+coverage.json|.+cobertura-coverage.xml|.+lcov.+")

// Reader reads and parses a coverage file.
type Reader interface {

	// Read reads a coverage report from the bytes.
	Read(src []byte) ([]*cover.Profile, error)

	// ReadFile reads a coverage report from the file path.
	ReadFile(string) ([]*cover.Profile, error)

	// ReadFrom reads a coverage report from the io.Reader.
	ReadFrom(io.Reader) ([]*cover.Profile, error)
}

// Register registers a reader associated with the sniff pattern.
func Register(match string, r Reader) {
	sniffers[r] = []byte(match)
}

// FromBytes reads the first 512 bytes of slice and returns a coverage
// report reader that is capable of parsing the coverage data.
func FromBytes(data []byte) (bool, Reader) {
	if len(data) > sniffLen {
		data = data[:sniffLen]
	}
	firstNonWS := 0
	for ; firstNonWS < len(data) && isWS(data[firstNonWS]); firstNonWS++ {
	}
	if firstNonWS >= len(data) {
		return false, nil
	}
	data = data[firstNonWS:]
	for r, sig := range sniffers {
		if bytes.HasPrefix(data, sig) {
			return true, r
		}
	}
	return false, nil
}

// FromFile reads the first 512 bytes of the file and returns a coverage
// report reader that is capable of parsing the coverage file.
func FromFile(path string) (bool, Reader) {
	f, err := os.Open(path)
	if err != nil {
		return false, nil
	}
	r := io.LimitReader(f, sniffLen)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return false, nil
	}
	return FromBytes(b)
}

// IsMatch returns true if the coverage file path matches a regular expression
// of known patterns for coverage files.
func IsMatch(path string) bool {
	return includes.MatchString(path)
}

// helper function returns true if the byte is whitespace
func isWS(b byte) bool {
	switch b {
	case '\t', '\n', '\x0c', '\r', ' ':
		return true
	}
	return false
}
