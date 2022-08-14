package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomInt(t *testing.T) {
	x := GenerateRandomInt(3)

	assert.LessOrEqual(t, x, 3)
}

func TestGenerateRandomString(t *testing.T) {
	x := GenerateRandomString(7)

	assert.Len(t, x, 7)
}
