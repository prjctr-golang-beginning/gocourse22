package kmtx

import (
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

func newTokenFromClaims(cs jwt.MapClaims) (kmtxJwtToken, error) {
	token := kmtxJwtToken{}
	for key, val := range cs {
		switch key {
		case `iat`:
			token.Iat = int64(val.(float64))
		case `exp`:
			token.Exp = int64(val.(float64))
		case `roles`:
			token.Roles = val.([]any)
		case `email`:
			token.Email = val.(string)
		case `permissions`:
			token.Permissions = val.([]any)
		default:
			log.Printf(`undefined key <%s> with value %v in kmtx token found`, key, val)
		}
	}

	return token, nil
}

type kmtxJwtToken struct {
	// Iat is the time when the token was issued
	Iat int64 `json:"iat"`

	// Exp is the time when the token expires
	Exp int64 `json:"exp"`

	// Email is the email of the user
	Email string `json:"email"`

	// KMTX roles
	Roles []any `json:"roles"`

	// KMTX permissions
	Permissions []any `json:"permissions"`
}

func (t *kmtxJwtToken) IsExpired() bool {
	return time.Now().UTC().Unix() >= t.Exp
}
