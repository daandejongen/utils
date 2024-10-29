package chainable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	assert.True(t, Chainable[int]{2, 5, 3, 7}.Contains(3))
	assert.False(t, Chainable[int]{2, 5, 3, 7}.Contains(8))
}
