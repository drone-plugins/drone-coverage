package coverage

import "testing"

func TestErrors(t *testing.T) {
	_, err := PathPrefix("/d/e/f", "/a/b/c")
	if err == nil {
		t.Errorf("Expect error with two different abolute paths")
	}

	_, err = PathPrefix("d/e/f", "/a/b/c")
	if err == nil {
		t.Errorf("Expect error with a relative path not in absolute")
	}
}

func TestPrefix(t *testing.T) {
	abs, err := PathPrefix("/a/b/c/d", "/a/b/c/")
	if err != nil && abs != "/a/b/c/" {
		t.Errorf("Expect prefix from two absolute paths %s", abs)
	}

	rel, err := PathPrefix("b/c/d/e", "/a/b/c/")
	if err != nil && rel != "b/c/" {
		t.Errorf("Expect prefix from a relative path %s", rel)
	}
}
