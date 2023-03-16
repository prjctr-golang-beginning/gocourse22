package local

func newUser(token authToken) *User {
	return &User{
		name:        token.username,
		token:       token,
		permissions: parseLocalPermissions(token.permissions),
	}
}

type User struct {
	name        string
	token       authToken
	permissions []string
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
	return `local`
}
