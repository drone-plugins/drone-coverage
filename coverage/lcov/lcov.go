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

	profiles := []*cover.Profile{}
	profile := &cover.Profile{}
	for s.Scan() {
		line := s.Text()

		switch {
		case strings.HasPrefix(line, "SF:"):
			profile = &cover.Profile{}
			profile.FileName = line[3:]
			profile.Mode = "set"
			profiles = append(profiles, profile)
		case strings.HasPrefix(line, "DA:"):
			parts := strings.Split(line[3:], ",")
			line, _ := strconv.Atoi(parts[0])
			count, _ := strconv.Atoi(parts[1])

			block := cover.ProfileBlock{}
			block.NumStmt = 1
			block.StartLine = line
			block.EndLine = line
			block.Count = count

			// TODO: not sure how to get this data yet
			block.StartCol = 0
			block.EndCol = 0

			profile.Blocks = append(profile.Blocks, block)
		default:
			continue
		}
	}

	return profiles, nil
}
