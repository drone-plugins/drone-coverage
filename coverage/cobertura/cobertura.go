package cobertura

import (
	"encoding/xml"
	"io"
	"io/ioutil"

	"github.com/drone-plugins/drone-coverage/coverage"

	"golang.org/x/tools/cover"
)

type cobertura struct {
	XMLName xml.Name `xml:"coverage"`
	Classes []class  `xml:"packages>package>classes>class"`
}

type class struct {
	Filename string `xml:"filename,attr"`
	Lines    []line `xml:"lines>line"`
}

type line struct {
	Number int `xml:"number,attr"`
	Hits   int `xml:"hits,attr"`
}

func init() {
	coverage.Register(`<?xml version="1.0">
<!DOCTYPE coverage SYSTEM "http://cobertura`, New())
	coverage.Register(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE coverage SYSTEM "http://cobertura`, New())
}

type reader struct {
}

// New returns a new Reader for reading and parsing a Cobertura report.
func New() coverage.Reader {
	return new(reader)
}

func (r *reader) Read(src []byte) ([]*cover.Profile, error) {

	cov, err := parse(src)

	if err != nil {
		return nil, err
	}

	var profiles []*cover.Profile

	for i, cls := range cov.Classes {

		var blocks []cover.ProfileBlock

		for _, line := range cls.Lines {
			blocks = append(blocks, cover.ProfileBlock{
				Count:     line.Hits,
				StartLine: line.Number,
				EndLine:   line.Number,
				NumStmt:   1,
			})
		}

		isNewFile := i == 0 || cov.Classes[i-1].Filename != cls.Filename

		if isNewFile {
			prof := &cover.Profile{}
			prof.FileName = cls.Filename
			prof.Mode = "set"
			prof.Blocks = append(prof.Blocks, blocks...)
			profiles = append(profiles, prof)
		} else {
			profiles[i-1].Blocks = append(profiles[i-1].Blocks, blocks...)
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

func parse(src []byte) (c cobertura, err error) {
	return c, xml.Unmarshal(src, &c)
}
