package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	a = &Key{"A"}
	b = &Key{"B"}
)

func Test_put(t *testing.T) {
	bst := InitBST()
	bst.Put(a, "0")
	bst.Put(b, "1")

	assert.Equal(t, bst.Get(a), "0")
	assert.Equal(t, bst.Get(b), "1")
}
