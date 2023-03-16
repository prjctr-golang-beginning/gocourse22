package kmtx

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

var testJWTToken = `eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpYXQiOjE2NzI4NDE2MzksImV4cCI6MTY3Mjg1MjQzOSwicm9sZXMiOltdLCJlbWFpbCI6Im0ubW9yb3pvdkBhdXRvZG9jLmV1IiwicGVybWlzc2lvbnMiOlsibWRtOnByb2R1Y3RzOnByb2R1Y3Q6dmlldyIsIm1kbTpicmFuZHM6YnJhbmQ6dmlldyIsIm1kbTpjb3VudHJpZXM6Y291bnRyeTp2aWV3IiwibWRtOmxhbmd1YWdlczpsYW5ndWFnZTp2aWV3IiwibWRtOmNyaXRlcmlhczpjcml0ZXJpYTp2aWV3IiwibWRtOmdlbmVyaWNzOmdlbmVyaWM6dmlldyIsIm1kbTpzdXBwbGllcnM6c3VwcGxpZXI6dmlldyIsIm1kbTpwcm9qZWN0czpwcm9qZWN0OnZpZXciLCJtZG06bWFudWZhY3R1cmVyczptYW51ZmFjdHVyZXI6dmlldyIsIm1kbTp2YWxpZGF0aW9uczp2YWxpZGF0aW9uOnZpZXciLCJtZG06cHJvZHVjdHM6cHJvZHVjdDp2aWV3IiwibWRtOnByb2R1Y3RzOnByb2R1Y3Q6ZWRpdCIsIm1kbTpwcm9kdWN0czpwcm9kdWN0OmNyZWF0ZSIsIm1kbTpwcm9kdWN0czpwcm9kdWN0OmRlbGV0ZSIsIm1kbTpicmFuZHM6YnJhbmQ6dmlldyIsIm1kbTpicmFuZHM6YnJhbmQ6Y3JlYXRlIiwibWRtOmJyYW5kczpicmFuZDpkZWxldGUiLCJtZG06YnJhbmRzOmJyYW5kOmVkaXQiLCJtZG06Y291bnRyaWVzOmNvdW50cnk6dmlldyIsIm1kbTpsYW5ndWFnZXM6bGFuZ3VhZ2U6dmlldyIsIm1kbTpjcml0ZXJpYXM6Y3JpdGVyaWE6dmlldyIsIm1kbTpjcml0ZXJpYXM6Y3JpdGVyaWE6ZWRpdCIsIm1kbTpjcml0ZXJpYXM6Y3JpdGVyaWE6ZGVsZXRlIiwibWRtOmNyaXRlcmlhczpjcml0ZXJpYTpjcmVhdGUiLCJtZG06Z2VuZXJpY3M6Z2VuZXJpYzp2aWV3IiwibWRtOmdlbmVyaWNzOmdlbmVyaWM6Y3JlYXRlIiwibWRtOmdlbmVyaWNzOmdlbmVyaWM6ZWRpdCIsIm1kbTpnZW5lcmljczpnZW5lcmljOmRlbGV0ZSIsIm1kbTpzdXBwbGllcnM6c3VwcGxpZXI6dmlldyIsIm1kbTpzdXBwbGllcnM6c3VwcGxpZXI6Y3JlYXRlIiwibWRtOnN1cHBsaWVyczpzdXBwbGllcjplZGl0IiwibWRtOnN1cHBsaWVyczpzdXBwbGllcjpkZWxldGUiLCJtZG06cHJvamVjdHM6cHJvamVjdDp2aWV3IiwibWRtOnByb2plY3RzOnByb2plY3Q6Y3JlYXRlIiwibWRtOnByb2plY3RzOnByb2plY3Q6ZWRpdCIsIm1kbTpwcm9qZWN0czpwcm9qZWN0OmRlbGV0ZSIsIm1kbTptYW51ZmFjdHVyZXJzOm1hbnVmYWN0dXJlcjp2aWV3IiwibWRtOm1hbnVmYWN0dXJlcnM6bWFudWZhY3R1cmVyOmNyZWF0ZSIsIm1kbTptYW51ZmFjdHVyZXJzOm1hbnVmYWN0dXJlcjplZGl0IiwibWRtOm1hbnVmYWN0dXJlcnM6bWFudWZhY3R1cmVyOmRlbGV0ZSIsIm1kbTp2YWxpZGF0aW9uczp2YWxpZGF0aW9uOnZpZXciLCJtZG06dmFsaWRhdGlvbnM6dmFsaWRhdGlvbjpjcmVhdGUiLCJtZG06dmFsaWRhdGlvbnM6dmFsaWRhdGlvbjplZGl0IiwibWRtOnZhbGlkYXRpb25zOnZhbGlkYXRpb246ZGVsZXRlIl19.CM7G9TfRDKhCY6Swl6sSgGzeuC9t8zMIJZ1HVURiz1gNHV29IB3QVuasYMHTzyoIqPFfgL3jWhXRV2bWUjyYlucuRDxwSDQn1k_4Y38FY-vQW8uEg8VdkK_qsw1IgC9bpOR3FXxjDXtWasz1DI4c08DDip1zMu3Ml6uVfzNpl73pGQAWygHocNcO3i6lubIUA8ofB_vwd6iG0Bj6LKdyQpAEXEdi9tNUhVpY_E5YuBMoFU5cmiiYKAv329DzjHZ8sZoGXkpcAZR5kMMBgZomGBl21btxxCTjYVgfNkwZ6hbX3mpfJ1iL0QWUtTo9a_VPsje1bRz94-PvDHAnK_co7sLzZ0Wy39YIV5acg69hzOeTHKaQaKTiKb_hwgJWswUPfYHW5B-jPF2PLhidbPwiBkrQ6Swk5nNsZBLx5cinSGZnzOVELMaZuncfgYYCoGkh4ZciVfMixf4lkOXPpKUTphygxGsVBOi7HUdBC92QJuZQh_rtJbOD2HQNxqDMwI_bMFXEFQo5pLnawRDo-dLP3vZhMVXvp_rRlb-ijTw4YTEaz_G_71VgKIWK6weTgdTl6nUOo867AjEqENhzqDL5syYXuIFGm63-8wf_rY1nn_0pNcyNwn179Rk9l-ESE2jdmRXr34eV-aCDL8IzZp0kK4VIWQduMF36ootRTcoFqSE`
var testAuthConfPubKey = `-----BEGIN PUBLIC KEY-----
...
-----END PUBLIC KEY-----`

func TestNewKMTXBearerAuthenticator(t *testing.T) {
	ba, err := NewKMTXBearerAuthenticator(AuthConfig{testAuthConfPubKey})

	assert.NoError(t, err)
	assert.NotEqualf(t, nil, ba.key, `key is nil`)
}

func TestBearerAuthor_Authenticate(t *testing.T) {
	ba, _ := NewKMTXBearerAuthenticator(AuthConfig{testAuthConfPubKey})

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
			name:      `Incorrect bearer auth header`,
			header:    `Bearer some-text-` + testJWTToken,
			errPrefix: `token error`,
			userIsNil: true,
		},
		//{ // TODO fix problem with key parsing
		//	name:      `Correct bearer auth header`,
		//	header:    `Bearer ` + testJWTToken,
		//	errPrefix: ``,
		//	userIsNil: false,
		//},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := &http.Request{Header: make(http.Header)}
			r.Header.Add(`Authorization`, c.header)
			user, err := ba.Authenticate(r)

			assert.True(t, strings.HasPrefix(err.Error(), c.errPrefix))
			assert.Equal(t, user == nil, c.userIsNil)
		})
	}
}
