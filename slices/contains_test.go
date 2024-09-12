package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	assert.True(t, Contains([]int{2, 5, 3, 7}, 3))
	assert.False(t, Contains([]int{2, 5, 3, 7}, 8))
}
