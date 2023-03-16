package local

type authToken struct {
	username    string
	expired     bool
	permissions []string
}

func (t *authToken) IsExpired() bool {
	return t.expired
}
