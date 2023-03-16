package kmtx

import "tests/pkg/auth/user"

func newUser(token kmtxJwtToken) user.User {
	return &User{
		name:        token.Email,
		token:       token,
		permissions: token.Permissions,
	}
}

type User struct {
	name        string
	token       kmtxJwtToken
	permissions []any
}

func (u *User) Name() string {
	return u.name
}

func (u *User) IsTokenExpired() bool {
	return u.token.IsExpired()
}

func (u *User) Permissions() []any {
	return u.permissions
}

func (u *User) Provider() string {
	return `kmtx`
}
