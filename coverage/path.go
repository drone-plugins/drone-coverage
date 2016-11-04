package coverage

import "fmt"

// PathPrefix finds the prefix relative to the base with the curr value.
// It will search the base to find the commonality or return an error if there
// is none.
func PathPrefix(curr string, base string) (string, error) {
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
