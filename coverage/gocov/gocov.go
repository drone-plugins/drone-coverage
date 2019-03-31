package gocov

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"

	"github.com/drone-plugins/drone-coverage/coverage"
	"golang.org/x/tools/cover"
)

func init() {
	coverage.Register("mode:", New())
}

type reader struct{}

// New returns a new Reader for reading and parsing a Go coverage report.
func New() coverage.Reader {
	return new(reader)
}

func (r *reader) Read(src []byte) ([]*cover.Profile, error) {
	buf := bytes.NewBuffer(src)
	return r.ReadProfiles(buf)
}

func (r *reader) ReadFile(path string) ([]*cover.Profile, error) {
	return cover.ParseProfiles(path)
}

func (r *reader) ReadProfiles(src io.Reader) ([]*cover.Profile, error) {
	file, err := ioutil.TempFile(os.TempDir(), "cover_file_")
	if err != nil {
		return nil, err
	}
	defer os.Remove(file.Name())
	if _, err := io.Copy(file, src); err != nil {
		return nil, err
	}
	return r.ReadFile(file.Name())
}
