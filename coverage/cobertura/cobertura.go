package cobertura

import (
	"encoding/xml"
	"io"
	"io/ioutil"

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

	type Line struct {
		Number int `xml:"number,attr"`
		Hits   int `xml:"hits,attr"`
	}

	type Method struct {
		Name       string `xml:"name,attr"`
		Signature  string `xml:"signature,attr"`
		LineRate   int    `xml:"line-rate,attr"`
		BranchRate int    `xml:"branch-rate,attr"`
		Lines      []Line `xml:"lines>line"`
	}

	type Class struct {
		Name       string   `xml:"name,attr"`
		Filename   string   `xml:"filename,attr"`
		LineRate   int      `xml:"line-rate,attr"`
		BranchRate int      `xml:"branch-rate,attr"`
		Complexity int      `xml:"complexity,attr"`
		Methods    []Method `xml:"methods>method"`
	}

	type Package struct {
		Name       string  `xml:"name,attr"`
		LineRate   int     `xml:"line-rate,attr"`
		BranchRate int     `xml:"branch-rate,attr"`
		Complexity int     `xml:"complexity,attr"`
		Classes    []Class `xml:"classes>class"`
	}

	type Coverage struct {
		XMLName    xml.Name  `xml:"coverage"`
		LineRate   int       `xml:"line-rate,attr"`
		BranchRate int       `xml:"branch-rate,attr"`
		Version    string    `xml:"version,attr"`
		Timestamp  int64     `xml:"timestamp,attr"`
		Packages   []Package `xml:"packages>package"`
	}

	cov := Coverage{}
	if err := xml.Unmarshal(src, &cov); err != nil {
		return nil, err
	}

	var profiles = []*cover.Profile{}

	for _, pkg := range cov.Packages {
		for i, cls := range pkg.Classes {
			profile := &cover.Profile{}
			profile.FileName = cls.Filename
			profile.Mode = "set"

			for _, meth := range cls.Methods {

				for _, line := range meth.Lines {
					profile.Blocks = append(profile.Blocks,
						cover.ProfileBlock{
							Count:     line.Hits,
							StartLine: line.Number,
							EndLine:   line.Number,
							NumStmt:   1,
						})
				}
			}

			if i == 0 || pkg.Classes[i-1].Filename != cls.Filename {
				profiles = append(profiles, profile)
			} else {
				for _, block := range profile.Blocks {
					profiles[i-1].Blocks = append(profiles[i-1].Blocks, block)
				}
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
