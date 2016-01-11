package client

// Client to access the remote APIs.
type Client interface {
	Token(string) (*Token, error)
	Repo(string) (*Repo, error)
	Activate(string) (*Repo, error)
	Deactivate(string) error
	Submit(string, *Build, *Report) (*Build, error)
}
