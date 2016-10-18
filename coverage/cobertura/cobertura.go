package cobertura

import (
	"bytes"
	"encoding/xml"
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

type Coverage struct {
	XMLName    xml.Name  `xml:"coverage"`
	LineRate   int       `xml:"line-rate,attr"`
	BranchRate int       `xml:"branch-rate,attr"`
	Version    string    `xml:"version,attr"`
	Timestamp  int64     `xml:"timestamp,attr"`
	Packages   []Package `xml:"packages>package"`
}

type Package struct {
	Name       string  `xml:"name,attr"`
	LineRate   int     `xml:"line-rate,attr"`
	BranchRate int     `xml:"branch-rate,attr"`
	Complexity int     `xml:"complexity,attr"`
	Classes    []Class `xml:"classes>class"`
}

type Class struct {
	Name     string `xml:"name,attr"`
	Filename string `xml:"filename,attr"`
}

func (r *reader) parseXML(src []byte) (c Coverage, err error) {
	return c, xml.Unmarshal(src, &c)
}
