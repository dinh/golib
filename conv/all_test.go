package conv

import (
	"github.com/funkygao/assert"
	"testing"
)

func TestConsts(t *testing.T) {
	assert.Equal(t, int32(48), ascii_0)
	assert.Equal(t, int32(57), ascii_9)
}

func TestParseInt(t *testing.T) {
	n, e := ParseInt([]byte("abc"))
	assert.Equal(t, ErrInvalidFormat, e)
	assert.Equal(t, -1, n)

	n, e = ParseInt([]byte("125"))
	assert.Equal(t, 125, n)
	assert.Equal(t, nil, e)

	// float not supported
	n, e = ParseInt([]byte("12.5"))
	assert.Equal(t, ErrInvalidFormat, e)
	assert.Equal(t, -1, n)

	// edge case
	n, _ = ParseInt([]byte("012"))
	assert.Equal(t, 12, n)
	n, _ = ParseInt([]byte("000012"))
	assert.Equal(t, 12, n)
	n, _ = ParseInt([]byte("0129"))
	assert.Equal(t, 129, n)
	n, _ = ParseInt([]byte("99999"))
	assert.Equal(t, 99999, n)
}
