package client

// Token represents a user oauth2 token.
type Token struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
	Expires int64  `json:"expires_in"`
}

// User represents a registered user.
type User struct {
	Login  string `json:"login"`
	Avatar string `json:"avatar_url"`
}

// Repo represents a remote version control repository
// for which coverage data has been collected.
type Repo struct {
	Owner    string  `json:"owner"`
	Name     string  `json:"name"`
	Slug     string  `json:"slug"`
	Link     string  `json:"link"`
	Branch   string  `json:"branch"`
	Avatar   string  `json:"avatar_url"`
	Private  bool    `json:"private"`
	Coverage float64 `json:"coverage_percent"`
	Delta    float64 `json:"coverage_changed"`
	Covered  int64   `json:"lines_covered"`
	Lines    int64   `json:"lines_total"`
}

// Build represents a build in an external continuous integration
// server for which coverage data has been collected.
type Build struct {
	Number    int     `json:"number"`
	Event     string  `json:"event"`
	Commit    string  `json:"commit"`
	Branch    string  `json:"branch"`
	Ref       string  `json:"ref"`
	Refspec   string  `json:"refspec"`
	Message   string  `json:"message"`
	Author    string  `json:"author"`
	Avatar    string  `json:"author_avatar"`
	Timestamp int64   `json:"timestamp"`
	Link      string  `json:"link_url"`
	Coverage  float64 `json:"coverage_percent"`
	Changed   float64 `json:"coverage_changed"`
	Covered   int64   `json:"lines_covered"`
	Lines     int64   `json:"lines_total"`
}

// File represents a source file from your repository that
// includes coverage data per line.
type File struct {
	FileName string  `json:"filename"`
	Coverage float64 `json:"coverage_percent"`
	Changed  float64 `json:"coverage_changed"`
	Covered  int64   `json:"lines_covered"`
	Lines    int64   `json:"lines_total"`
	Mode     string  `json:"coverage_mode"`

	Blocks []*Block `json:"blocks"`
}

// Block represents a block of code in a source code file
// and includes coverage details.
type Block struct {
	StartLine int `json:"start_line"`
	StartCol  int `json:"start_col"`
	EndLine   int `json:"end_line"`
	EndCol    int `json:"end_col"`
	NumStmt   int `json:"num_stmt"`
	Count     int `json:"count"`
}

// Report represents a code coverage report.
type Report struct {
	Coverage float64 `json:"coverage_percent"`
	Changed  float64 `json:"coverage_changed"`
	Covered  int64   `json:"lines_covered"`
	Lines    int64   `json:"lines_total"`
	Files    []*File `json:"files"`
}
