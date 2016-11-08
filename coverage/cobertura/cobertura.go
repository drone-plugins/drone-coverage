package cobertura

import (
	"bufio"
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"

	"github.com/drone-plugins/drone-coverage/coverage"

	"golang.org/x/tools/cover"
)

type cobertura struct {
	XMLName xml.Name `xml:"coverage"`
	Classes []class  `xml:"packages>package>classes>class"`
}

type class struct {
	FileName string `xml:"filename,attr"`
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
	coverage.Register(`<?xml version="1.0" ?>
<coverage`, New())
	coverage.Register(`<?xml version="1.0" encoding="UTF-8"?>
<coverage`, New())
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
		cols := calculateCols(cls.FileName)

		for _, line := range cls.Lines {
			block := cover.ProfileBlock{
				Count:     line.Hits,
				StartLine: line.Number,
				EndLine:   line.Number,
				NumStmt:   1,
			}

			if len(cols) != 0 {
				block.StartCol = 1
				block.EndCol = cols[line.Number-1]
			}

			blocks = append(blocks, block)
		}

		isNewFile := i == 0 || cov.Classes[i-1].FileName != cls.FileName

		if isNewFile {
			profile := &cover.Profile{}

			profile.FileName = cls.FileName
			profile.Mode = "set"
			profile.Blocks = append(profile.Blocks, blocks...)
			profiles = append(profiles, profile)
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

// this is a helper function that is able to calculate the number of columns
func calculateCols(fileName string) []int {
	var lines []int

	f, err := os.Open(fileName)
	if err != nil {
		return lines
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cols := len(scanner.Bytes()) + 1
		lines = append(lines, cols)
	}
	return lines
}
