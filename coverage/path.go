package coverage

import (
	"fmt"
	"path"
	"strings"
)

func PathPrefix(curr string, base string) (string, error) {
	// Check for absolute paths first
	if path.IsAbs(curr) {
		if !strings.HasPrefix(curr, base) {
			return "", fmt.Errorf("Path %s not found in %s", curr, base)
		}

		return base, nil
	}

	cBytes := []byte(curr)
	bBytes := []byte(base)
	count := len(bBytes)

	// Search for the string
	for i, _ := range bBytes {
		for j, c := range cBytes {
			a := i + j

			if a == count {
				return curr[:j], nil
			}

			if c != bBytes[a] {
				break
			}
		}
	}

	return "", fmt.Errorf("Path %s not found in %s", curr, base)
}
