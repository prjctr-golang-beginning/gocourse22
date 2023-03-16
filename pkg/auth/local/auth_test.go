package local

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

var filename = `local_users_test.yml`

func TestNewLocalAuthenticator(t *testing.T) {
	la, err := NewLocalAuthenticator(AuthConfig{filename})

	assert.NoError(t, err)
	assert.NotEqualf(t, nil, la.storage.PermissionGroups, `PermissionGroups is nil`)
	assert.NotEqualf(t, nil, la.storage.Users, `Users is nil`)
}

func TestAuthor_Authenticate(t *testing.T) {
	ba, _ := NewLocalAuthenticator(AuthConfig{filename})

	cases := []struct {
		name      string
		header    string
		errPrefix string
		userIsNil bool
	}{
		{
			name:      `No auth header`,
			header:    ``,
			errPrefix: authErrHNF.Error(),
			userIsNil: true,
		},
		{
			name:      `Only local`,
			header:    `not-local`,
			errPrefix: authErrOL.Error(),
			userIsNil: true,
		},
		{
			name:      `Empty user`,
			header:    `local username=`,
			errPrefix: authErrEU.Error(),
			userIsNil: true,
		},
		{
			name:      `Incorrect user`,
			header:    `local username=someuser`,
			errPrefix: authErrUNF.Error(),
			userIsNil: true,
		},
		{
			name:      `Correct user`,
			header:    `local username=test_user`,
			errPrefix: ``,
			userIsNil: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := &http.Request{Header: make(http.Header)}
			r.Header.Add(`Authorization`, c.header)
			user, err := ba.Authenticate(r)

			if err != nil {
				assert.True(t, strings.HasPrefix(err.Error(), c.errPrefix))
			}
			assert.Equal(t, user == nil, c.userIsNil)
		})
	}
}
