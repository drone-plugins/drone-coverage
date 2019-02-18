package jacoco

import (
	"io"

	"github.com/drone-plugins/drone-coverage/coverage"
	"golang.org/x/tools/cover"
)

func init() {
	coverage.Register(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<!DOCTYPE report PUBLIC "-//JACOCO`, New())
}

type reader struct{}

// New returns a new Reader for reading and parsing a Jacoco report.
func New() coverage.Reader {
	return new(reader)
}

func (r *reader) Read(src []byte) ([]*cover.Profile, error) {
	return nil, nil
}

func (r *reader) ReadFile(path string) ([]*cover.Profile, error) {
	return nil, nil
}

func (r *reader) ReadProfiles(src io.Reader) ([]*cover.Profile, error) {
	return nil, nil
}
