package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasDirectDiff(t *testing.T) {
	assert.False(t, HasDirectDiff([]string{`f1`, `f2`, `f3`}, []string{`f1`, `f2`, `f3`, `f4`, `f5`}))
	assert.True(t, HasDirectDiff([]string{`f1`, `f2`, `f3`, `f7`}, []string{`f1`, `f2`, `f3`, `f4`, `f5`}))
	assert.True(t, HasDirectDiff([]string{`f1`, `f0`, `f3`}, []string{`f1`, `f2`, `f3`, `f4`, `f5`}))
	assert.False(t, HasDirectDiff([]string{`f1`, `f0`, `f3`}, []string{`f0`, `f1`, `f3`, `f4`, `f5`}))
	assert.False(t, HasDirectDiff([]string{}, []string{`f0`, `f1`}))
	assert.True(t, HasDirectDiff([]string{`f0`, `f1`}, []string{}))
}
