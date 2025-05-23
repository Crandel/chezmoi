//go:build unix

package chezmoi

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/chezmoi/internal/chezmoitest"
)

func TestUmask(t *testing.T) {
	assert.Equal(t, chezmoitest.Umask, Umask)
}
