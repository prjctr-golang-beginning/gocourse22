package user

type User interface {
	Name() string
	IsTokenExpired() bool
	Permissions() []any
	Provider() string
}
