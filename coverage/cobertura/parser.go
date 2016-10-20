package cobertura

import "encoding/xml"

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
	Name       string   `xml:"name,attr"`
	Filename   string   `xml:"filename,attr"`
	LineRate   int      `xml:"line-rate,attr"`
	BranchRate int      `xml:"branch-rate,attr"`
	Complexity int      `xml:"complexity,attr"`
	Methods    []Method `xml:"methods>method"`
}

type Method struct {
	Name       string `xml:"name,attr"`
	Signature  string `xml:"signature,attr"`
	LineRate   int    `xml:"line-rate,attr"`
	BranchRate int    `xml:"branch-rate,attr"`
	Lines      []Line `xml:"lines>line"`
}

type Line struct {
	Number int `xml:"number,attr"`
	Hits   int `xml:"hits,attr"`
}

// Parse parsers a cobertura xml file to coverage struct
func Parse(src []byte) (c Coverage, err error) {
	return c, xml.Unmarshal(src, &c)
}
