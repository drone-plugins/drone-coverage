package cobertura

import (
	"encoding/xml"
	"io"
	"io/ioutil"

	"github.com/drone-plugins/drone-coverage/coverage"

	"golang.org/x/tools/cover"
)

type cobertura struct {
	XMLName  xml.Name `xml:"coverage"`
	Packages []pkg    `xml:"packages>package"`
}

type pkg struct {
	Classes []class `xml:"classes>class"`
}

type class struct {
	Filename string   `xml:"filename,attr"`
	Methods  []method `xml:"methods>method"`
}

type method struct {
	Lines []line `xml:"lines>line"`
}

type line struct {
	Number int `xml:"number,attr"`
	Hits   int `xml:"hits,attr"`
}

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

	cov, err := parse(src)

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

func parse(src []byte) (c cobertura, err error) {
	return c, xml.Unmarshal(src, &c)
}

func getBlocksFromMethods(methods []method) (blocks []cover.ProfileBlock) {
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
