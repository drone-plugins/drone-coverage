package cobertura

import (
	"bytes"
	"io"
	"os"

	"github.com/drone-plugins/drone-coverage/coverage"
	"golang.org/x/tools/cover"
)

func init() {
	coverage.Register(`<?xml version="1.0" ?>
<!DOCTYPE coverage SYSTEM "http://cobertura`, New())
}

type reader struct{}

// New returns a new Reader for reading and parsing a Cobertura report.
func New() coverage.Reader {
	return new(reader)
}

func (r *reader) Read(src []byte) ([]*cover.Profile, error) {
	buf := bytes.NewBuffer(src)
	return r.ReadFrom(buf)
}

func (r *reader) ReadFile(path string) ([]*cover.Profile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return r.ReadFrom(file)
}

func (r *reader) ReadFrom(src io.Reader) ([]*cover.Profile, error) {
	// all reading logic goes here!
	return nil, nil
}

type Cobertura struct {
	Packages []Package
}

type Package struct {
	Name string
}

func (r *reader) parseXML(src []byte) (Cobertura, error) {

	c := Cobertura{
		Packages: []Package{
			{Name: "/home/fbcbarbosa/Development/go/src/github.com/drone-plugins/drone-coverage/coverage/gocov/gocov.go"},
		},
	}

	return c, nil
}
