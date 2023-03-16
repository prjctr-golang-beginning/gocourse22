package auth

import (
	"reflect"
	"testing"
	"tests/pkg/auth/kmtx"
	"tests/pkg/auth/local"
)

func Test_GetAuthenticator(t *testing.T) {
	authConfig := getDefaultAuthConfig()

	authConfigWithDefault := getDefaultAuthConfig()
	authConfigWithDefault.Provider.Default = "local"

	authConfigWithNoAllowed := getDefaultAuthConfig()
	authConfigWithNoAllowed.Provider.Allowed = ""

	kmtxAuth, kmtxErr := kmtx.NewKMTXBearerAuthenticator(kmtx.AuthConfig{PubKey: authConfig.KMTX.PubKey})
	if kmtxErr != nil {
		t.Errorf("Error creating kmtx auth: %v", kmtxErr)
	}

	localAuth, localErr := local.NewLocalAuthenticator(local.AuthConfig{FileName: authConfig.Local.FileName})
	if localErr != nil {
		t.Errorf("Error creating local auth: %v", localErr)
	}

	type args struct {
		header string
		cfg    Config
	}
	tests := []struct {
		name       string
		args       args
		want       Authenticator
		wantErr    bool
		wantErrStr string
	}{
		{
			name: "error when no header is provided and no default auth provider is set",
			args: args{
				header: "",
				cfg:    authConfig,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: "unknown auth provider \"\"",
		},

		{
			name: "error when set unknown auth provider in header",
			args: args{
				header: "unknown",
				cfg:    authConfig,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: "unsupported authentication type \"unknown\"",
		},

		{
			name: "error when unallowed auth provider in header",
			args: args{
				header: "kmtx",
				cfg:    authConfigWithNoAllowed,
			},
			want:       nil,
			wantErr:    true,
			wantErrStr: "unsupported authentication type \"kmtx\"",
		},

		{
			name: "get kmtx authenticator when kmtx is set in header",
			args: args{
				header: "kmtx",
				cfg:    authConfig,
			},
			want:       kmtxAuth,
			wantErr:    false,
			wantErrStr: "",
		},

		{
			name: "get local authenticator when local is set in header",
			args: args{
				header: "local",
				cfg:    authConfig,
			},
			want:       localAuth,
			wantErr:    false,
			wantErrStr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAuthenticator(tt.args.header, tt.args.cfg)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("getAuthenticator() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err.Error() != tt.wantErrStr {
					t.Errorf("getAuthenticator() error = %v, wantErr %v", err, tt.wantErrStr)
					return
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAuthenticator() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func getDefaultAuthConfig() Config {
	return Config{
		Provider: struct {
			Allowed string `default:"local" validate:"required,check-allowed-providers"`
			Default string `default:"local" validate:"omitempty,oneof=kmtx aws local"`
		}{
			Allowed: "local,aws,kmtx",
			Default: "",
		},

		KMTX: struct {
			PubKey string `validate:"omitempty,min=1,valid-rsa-pub-key"`
		}{PubKey: "-----BEGIN PUBLIC KEY-----\nMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAw2DE/EQTJbR1rI047Pzb\nmzKu/t3MClLyl5YNhHBdWv6v6vUzEuQcvYapE4H+5OY48rbuZR0KqjJyMEdsBB0g\nQfAJuzwSHoLEdfed1ggd+ovFNVcFczbZUIK4Yb1d55oDz/Bi5VqYp8IjMtWNzWK4\nJ/ULbDR6tCEy6ppW1lWL0kTNk2RV/tKeK4xKQvlROcMHgnsLLtQ1UhikuX/URSuS\nAz6P0LpUgIyOXG5eZirEIGYpjp2lIKdpAL5kqh3Z8UiKKgQ0/Q2xLdRsIfwL0LtS\ngqNB2TsoPh9rAW9rVlvdsrYXL7eWPR+1Ih4YLrNR42kmSY3W0uezcY4B6EPkMOHl\ndMTq2LynAss/aykRChGJeL7gfkFvhS/KJ6+Hn5xgCdl6VrDDfXQ6vchihHT2YX8d\nGh1QREBbWPcqtjZQufPFpAPaUh/T6zk797biaCl85ylV0u+TCjewPH6MxjAO+WWF\nVOeeecV03CFwM5CRpGtyDH2GY0rvZ9eK11g/N4rbulc9vw1XACsIo61NQaHklMLI\nTE6g15m8pSg/+aufxD6FKQxkSxvllBh6dQ9yP56p+ZT/Wt78C21ZJ/KxmsXSyyLX\nlZVluG4z5Zgyvp6etXAGfnWUY0XGeuT3U8JgFxDvNhlg/Diri9lKDi2BjKlahQ02\nv8KYJSnC43FOssrP8/k1oS8CAwEAAQ==\n-----END PUBLIC KEY-----"},

		Local: struct {
			FileName string `validate:"omitempty,min=1"`
		}{FileName: "../../../data/local_users.yml"},
	}
}
