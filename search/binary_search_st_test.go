package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchST_IsEmpty(t *testing.T) {
	list := InitBinarySearchST(100)
	assert.Equal(t, list.IsEmpty(), true)
}

type key struct {
	k int
}

func (i *key) Compare(j BSSTI) int {
	_, ok := j.(*key)
	println(ok)
	return i.k - j.(*key).k
}

func TestBinarySearchST_Set(t *testing.T) {
	list := InitBinarySearchST(100)
	assert.Equal(t, list.IsEmpty(), true)

	list.Set(&key{1}, 123)
	assert.Equal(t, list.Size(), 1)
	assert.Equal(t, list.Get(&key{1}), 123)
}

func TestBinarySearchST_Get(t *testing.T) {
	list := InitBinarySearchST(100)
	assert.Equal(t, list.IsEmpty(), true)

	list.Set(&key{1222}, 123)
	list.Set(&key{12223}, 123)
	assert.Equal(t, list.Get(&key{1222}), 123)
	assert.Equal(t, list.Get(&key{12223}), 123)
}
