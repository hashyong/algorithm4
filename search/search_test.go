package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeqSearchST_Get(t *testing.T) {
	list := InitNode("my", 1)

	assert.Equal(t, list.Get("my"), 1)

}

func TestSeqSearchST_Set(t *testing.T) {
	list := InitNode("my", 1)
	list.Set("test", "test")
	list.Set("test1", "test1")

	assert.Equal(t, list.Get("test1"), "test1")
	assert.Equal(t, list.Get("test"), "test")
}

func TestSeqSearchST_Del(t *testing.T) {
	list := InitNode("my", 1)
	list.Set("test", "test")
	list.Set("test1", "test1")
	assert.Equal(t, list.Get("test1"), "test1")
	list.Del("test1")
	assert.Equal(t, list.Get("test1"), nil)
}

func TestSeqSearchST_Size(t *testing.T) {
	list := InitNode("my", 1)
	list.Set("test", "test")
	list.Set("test1", "test1")
	assert.Equal(t, list.Size(), 3)
}
