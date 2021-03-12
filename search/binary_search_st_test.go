package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchST_IsEmpty(t *testing.T) {
	list := InitBinarySearchST(100)
	assert.Equal(t, list.IsEmpty(), true)
}
