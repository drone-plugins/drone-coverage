package lcov

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/drone-plugins/drone-coverage/coverage"
	"golang.org/x/tools/cover"
)

func init() {
	coverage.Register("SF:", New())
	coverage.Register("TN:", New())
}

type reader struct{}

// New returns a new Reader for reading and parsing an LCOV report.
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
	buf := bufio.NewReader(src)
	s := bufio.NewScanner(buf)

	var profiles []*cover.Profile
	var cols []int

	profile := &cover.Profile{}
	for s.Scan() {
		line := s.Text()

		switch {
		case strings.HasPrefix(line, "SF:"):
			profile = &cover.Profile{}
			if strings.HasPrefix(line[3:], "./") {
				profile.FileName = strings.Replace(line[3:], "./", "/", 1)
			} else {
				profile.FileName = line[3:]
			}
			profile.Mode = "set"
			profiles = append(profiles, profile)

			// until I can think of a better way we need to know how many
			// columns are in each line of code.
			cols = calculateCols(profile)

		case strings.HasPrefix(line, "DA:"):
			parts := strings.Split(line[3:], ",")
			line, _ := strconv.Atoi(parts[0])
			count, _ := strconv.Atoi(parts[1])
			if count > 0 {
				count = 1
			}

			block := cover.ProfileBlock{}
			block.NumStmt = 1
			block.StartLine = line
			block.EndLine = line
			block.Count = count

			if len(cols) != 0 && line > 0 {
				block.StartCol = 1
				block.EndCol = cols[line]
			}

			profile.Blocks = append(profile.Blocks, block)
		default:
			continue
		}
	}

	return profiles, nil
}

// this is a helper function that is able to calculate the
func calculateCols(profile *cover.Profile) []int {
	var lines []int

	f, err := os.Open(profile.FileName)
	if err != nil {
		return lines
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cols := len(scanner.Bytes())
		lines = append(lines, cols)
	}
	return lines
}
