package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringContainsAnyKeyword(t *testing.T) {
	assert.True(t, StringContainsAnyKeyword("beer", "wine", "beer", "coke"))
	assert.False(t, StringContainsAnyKeyword("beer", "wine", "water", "coke"))
}