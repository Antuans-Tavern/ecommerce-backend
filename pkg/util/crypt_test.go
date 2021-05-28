package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAndCompare(t *testing.T) {
	hash := Hash("secret")

	assert.True(t, CompareHash(hash, "secret"))
}
