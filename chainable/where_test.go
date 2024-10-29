package chainable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhereIntIsPositive(t *testing.T) {
	isPositive := func(i int) bool { return i > 0 }
	input := Chainable[int]{0, 2, -4, -2, 1, 9, -2}
	expected := Chainable[int]{2, 1, 9}
	assert.True(t, input.Where(isPositive).Equals(expected))
}
