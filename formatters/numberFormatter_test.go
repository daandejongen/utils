package formatters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToTwoDigits(t *testing.T) {
	nf := NewNumberFormatter()
	assert.Equal(t, "00", nf.ToTwoDigits(0))
	assert.Equal(t, "02", nf.ToTwoDigits(2))
	assert.Equal(t, "14", nf.ToTwoDigits(14))
	assert.Equal(t, "24", nf.ToTwoDigits(1124))
}
