package cobertura

import (
	"io"
	"io/ioutil"

	"github.com/drone-plugins/drone-coverage/coverage"

	"golang.org/x/tools/cover"
)

func init() {
	coverage.Register(`<?xml version="1.0" ?>
<!DOCTYPE coverage SYSTEM "http://cobertura`, New())
}

type reader struct {
}

// New returns a new Reader for reading and parsing a Cobertura report.
func New() coverage.Reader {
	return new(reader)
}

func (r *reader) Read(src []byte) ([]*cover.Profile, error) {

	cov, err := Parse(src)

	if err != nil {
		return nil, err
	}

	var profiles []*cover.Profile

	for _, pkg := range cov.Packages {
		for i, cls := range pkg.Classes {

			blocks := getBlocksFromMethods(cls.Methods)

			if i == 0 || pkg.Classes[i-1].Filename != cls.Filename {
				prof := &cover.Profile{}
				prof.FileName = cls.Filename
				prof.Mode = "set"
				prof.Blocks = append(prof.Blocks, blocks...)
				profiles = append(profiles, prof)
			} else {
				profiles[i-1].Blocks = append(profiles[i-1].Blocks, blocks...)
			}
		}
	}

	return profiles, nil
}

func (r *reader) ReadFile(path string) ([]*cover.Profile, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return r.Read(data)
}

func (r *reader) ReadFrom(src io.Reader) ([]*cover.Profile, error) {
	data, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}
	return r.Read(data)
}

func getBlocksFromMethods(methods []Method) (blocks []cover.ProfileBlock) {
	for _, meth := range methods {
		for _, line := range meth.Lines {
			blocks = append(blocks, cover.ProfileBlock{
				Count:     line.Hits,
				StartLine: line.Number,
				EndLine:   line.Number,
				NumStmt:   1,
			})
		}
	}
	return
}
