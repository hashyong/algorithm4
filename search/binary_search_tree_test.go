package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	a = &Key{"A"}
	b = &Key{"B"}
	c = &Key{"C"}
	d = &Key{"D"}
	e = &Key{"E"}
)

func Test_put(t *testing.T) {
	bst := InitBST()
	bst.Put(a, "0")
	bst.Put(b, "1")

	assert.Equal(t, bst.Get(a), "0")
	assert.Equal(t, bst.Get(b), "1")
}

func TestBST_Range(t *testing.T) {
	bst := InitBST()
	bst.Put(b, "1")
	bst.Put(c, "1")
	bst.Put(a, "0")
	bst.Put(e, "1")
	bst.Put(d, "1")

	result := bst.Range(b, d)
	for _, item := range result {
		fmt.Println(item)
	}

	res := []BSSTI{b, c, d}
	assert.Equal(t, result, res)
}
