package auth

import (
	"fmt"
	"net/http"
	"sync"
	"tests/pkg/auth/kmtx"
	"tests/pkg/auth/local"
	"tests/pkg/auth/user"
)

var (
	kmtxBearerAuthor   *kmtx.BearerAuthor
	kmtxBearerAuthorMu sync.Once

	localBearerAuthor   *local.Author
	localBearerAuthorMu sync.Once
)

type Authenticator interface {
	Authenticate(r *http.Request) (user.User, error)
}

func GetAuthenticator(customProvider string, cfg Config) (Authenticator, error) {
	authProvider := cfg.Provider.Default
	if customProvider != "" {
		authProvider = customProvider
	}

	if !cfg.IsProviderAllowed(authProvider) {
		return nil, fmt.Errorf(`unsupported authentication type "%s"`, authProvider)
	}

	switch authProvider {
	case `kmtx`:
		var err error
		kmtxBearerAuthorMu.Do(func() {
			kmtxBearerAuthor, err = kmtx.NewKMTXBearerAuthenticator(kmtx.AuthConfig{
				PubKey: cfg.KMTX.PubKey,
			})
		})
		return kmtxBearerAuthor, err
	case `local`:
		var err error
		localBearerAuthorMu.Do(func() {
			localBearerAuthor, err = local.NewLocalAuthenticator(local.AuthConfig{
				FileName: cfg.Local.FileName,
			})
		})
		return localBearerAuthor, err
	default:
		return nil, fmt.Errorf(`unknown auth provider "%s"`, authProvider)
	}
}
