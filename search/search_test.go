package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testKey  = &Key{"123"}
	testKey1 = &Key{"test1"}
)

func TestSeqSearchST_Get(t *testing.T) {
	list := InitNode(testKey, 1)

	assert.Equal(t, list.Get("my"), 1)

}

func TestSeqSearchST_Set(t *testing.T) {
	list := InitNode(testKey, 1)
	list.Set(testKey, "test")
	list.Set(testKey1, "test1")

	assert.Equal(t, list.Get(testKey), "test1")
	assert.Equal(t, list.Get(testKey1), "test")
}

func TestSeqSearchST_Del(t *testing.T) {
	list := InitNode(testKey, 1)
	list.Set(testKey, "test")
	list.Set(testKey1, "test1")
	assert.Equal(t, list.Get(testKey1), "test1")
	list.Del("test1")
	assert.Equal(t, list.Get(testKey1), nil)
}

func TestSeqSearchST_Size(t *testing.T) {
	list := InitNode(testKey, 1)
	list.Set(testKey1, "test")
	assert.Equal(t, list.Size(), 2)
}
