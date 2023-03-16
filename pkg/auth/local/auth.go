package local

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"tests/pkg/auth/user"
)

var authErrHNF = errors.New(`authorization header not found`)
var authErrE = errors.New(`authorization header is empty`)
var authErrOL = errors.New(`authorization allowed only authentication type "local"`)
var authErrOIF = errors.New(`authorization header has invalid format`)
var authErrEU = errors.New(`authorization header contains empty username`)
var authErrUNF = errors.New(`user not found`)

func NewLocalAuthenticator(cfg AuthConfig) (*Author, error) {
	storage, err := parse(cfg.FileName)
	if err != nil {
		return nil, err
	}

	return &Author{storage: storage}, nil
}

type Author struct {
	storage *localStorage
}

func (a *Author) Authenticate(r *http.Request) (user.User, error) {
	header := r.Header.Get("Authorization")
	if header == `` {
		return nil, authErrHNF
	}

	parts := strings.Split(header, ` `)
	if len(parts) == 0 {
		return nil, authErrE
	}

	if strings.ToLower(parts[0]) != "local" {
		return nil, authErrOL
	}

	token := authToken{}
	for _, part := range parts[1:] {
		if strings.ContainsRune(part, '=') {
			p := strings.Split(part, `=`)
			if len(p) != 2 {
				return nil, authErrOIF
			}

			switch strings.ToLower(p[0]) {
			case `username`:
				token.username = p[1]
			case `expired`:
				token.expired = p[1] == `1`
			}
		} else {
			return nil, fmt.Errorf(`authorization header contains unexpected part %s`, part)
		}
	}

	if token.username == `` {
		return nil, authErrEU
	}

	if usr := a.storage.GetUser(token.username); usr == nil {
		return nil, authErrUNF
	} else {
		token.permissions = usr.Permissions
	}

	return newUser(token), nil
}
