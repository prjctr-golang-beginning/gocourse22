package kmtx

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

var parsedJWTToken jwt.MapClaims = map[string]any{
	"iat":   1672841639.0,
	"exp":   1672852439.0,
	"roles": []any{},
	"email": "my@mail.com",
	"permissions": []any{
		"project:products:product:view",
		"project:brands:brand:view",
		"project:countries:country:view",
		"project:languages:language:view",
	},
}

func TestNewTokenFromClaims(t *testing.T) {
	jwtToken, err := newTokenFromClaims(parsedJWTToken)

	assert.NoError(t, err)
	assert.NotEmpty(t, jwtToken.Permissions)
	assert.Empty(t, jwtToken.Roles)
}
