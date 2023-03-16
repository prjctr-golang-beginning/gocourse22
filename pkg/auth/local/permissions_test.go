package local

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLocalPermissions(t *testing.T) {
	cases := []struct {
		name   string
		perms  []string
		length int
	}{
		{
			name: `All is correct`,
			perms: []string{
				`product:product:create`,
				`brand:brand:view`,
				`brand:brand:create`,
			},
			length: 3,
		},
		{
			name: `Invalid permission`,
			perms: []string{
				`some-permission`,
				`brand:brand:create`,
			},
			length: 1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := parseLocalPermissions(c.perms)
			assert.Equal(t, c.length, len(res))
		})
	}
}
