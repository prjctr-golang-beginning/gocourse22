package kmtx

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"tests/pkg/auth/user"
)

var authErrHNF = errors.New(`authorization header not found`)

func NewKMTXBearerAuthenticator(cfg AuthConfig) (*BearerAuthor, error) {
	// init kmtx authenticator
	key, parseErr := jwt.ParseRSAPublicKeyFromPEM([]byte(cfg.PubKey))
	if parseErr != nil {
		return nil, parseErr
	}

	return &BearerAuthor{key: key}, nil
}

type BearerAuthor struct {
	key *rsa.PublicKey
}

func (a *BearerAuthor) Authenticate(r *http.Request) (user.User, error) {
	header := r.Header.Get("Authorization")
	if header == `` {
		return nil, authErrHNF
	}

	claims := jwt.MapClaims{}
	jwtToken, parseErr := jwt.ParseWithClaims(
		strings.TrimPrefix(header, `Bearer `),
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return a.key, nil
		},
	)

	if parseErr != nil {
		return nil, fmt.Errorf(`token error: %v`, parseErr)
	}

	if !jwtToken.Valid {
		return nil, errors.New(`invalid token`)
	}

	token, tokenErr := newTokenFromClaims(claims)
	if tokenErr != nil {
		return nil, tokenErr
	}

	return newUser(token), nil
}
