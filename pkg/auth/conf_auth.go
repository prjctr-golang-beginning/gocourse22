package auth

import (
	"fmt"
	"strings"
)

type Config struct {
	Provider struct {
		Allowed string `default:"local" validate:"required,check-allowed-providers"`
		Default string `default:"local" validate:"omitempty,oneof=kmtx aws local"`
	}
	KMTX struct {
		PubKey string `validate:"omitempty,min=1,valid-rsa-pub-key"`
	}
	Local struct {
		FileName string `validate:"omitempty,min=1"`
	}
}

func (a *Config) IsProviderAllowed(provider string) bool {
	return strings.Contains(a.Provider.Allowed, provider)
}

func (a *Config) String() string {
	kmtxPubKeyExists := len(a.KMTX.PubKey) > 0

	lines := []string{
		fmt.Sprintf("  Provider:\n    Default: %v\n    Allowed: %v", a.Provider.Default, strings.Split(a.Provider.Allowed, ",")),
		fmt.Sprintf("  KMTX:\n    Verified public key: %v", kmtxPubKeyExists),
		fmt.Sprintf("  Local:\n    Database path: %v", a.Local.FileName),
	}

	return strings.Join(lines, "\n")
}
