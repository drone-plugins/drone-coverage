package cobertura

import (
	"io"

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
	return nil, nil
}

func (r *reader) ReadFile(path string) ([]*cover.Profile, error) {
	return nil, nil
}

func (r *reader) ReadFrom(src io.Reader) ([]*cover.Profile, error) {
	return nil, nil
}
