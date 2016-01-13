package coverage

import "testing"

func TestFromBytes(t *testing.T) {
	Register("mode:", nil)

	// test empty string
	ok, _ := FromBytes([]byte(""))
	if ok {
		t.Errorf("Expect bytes to not match a reader")
	}

	// test whitespace that equals the pattern length
	ok, _ = FromBytes([]byte("    "))
	if ok {
		t.Errorf("Expect bytes to not match a reader")
	}

	// test whitespace greater than the pattern length
	ok, _ = FromBytes([]byte("     "))
	if ok {
		t.Errorf("Expect bytes to not match a reader")
	}

	// test a matching pattern
	ok, _ = FromBytes([]byte("mode:"))
	if !ok {
		t.Errorf("Expect bytes to match a reader")
	}

	// test whitespace prefixing a matching pattern
	ok, _ = FromBytes([]byte("  mode:"))
	if !ok {
		t.Errorf("Expect bytes to match a reader even with whitespace")
	}
}

func TestIsMatch(t *testing.T) {
	var tests = map[string]bool{
		"/drone/src/github.com/octocat/hello-world/coverage.out":           true,
		"/drone/src/github.com/octocat/hello-world/coverage.stdout":        false,
		"/drone/src/github.com/octocat/hello-world/clover.xml":             true,
		"/drone/src/github.com/octocat/hello-world/coverage.json":          true,
		"/drone/src/github.com/octocat/hello-world/cobertura-coverage.xml": true,
		"/drone/src/github.com/octocat/hello-world/lcov.info":              true,
		"/drone/src/github.com/octocat/hello-world/lcov.txt":               true,
	}

	for path, match := range tests {
		if IsMatch(path) != match {
			t.Errorf("Expected filepath match %v for %s", match, path)
		}
	}
}
