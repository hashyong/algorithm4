package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

func TestGraph_Graph(t *testing.T) {
	g := New().Graph(10)
	fmt.Println(toString(g))
}

func TestGraph_GraphIn(t *testing.T) {
	g := New().GraphIn("test.json")
	fmt.Println(toString(g))
}

func TestGraph_E(t *testing.T) {
	g := New().Graph(10)
	assert.Equal(t, g.E(), 0)
}

func TestGraph_V(t *testing.T) {
	v := 10
	g := New().Graph(v)
	assert.Equal(t, g.V(), v)
}

func TestGraph_AddEdge(t *testing.T) {
	v := 3
	g := New().Graph(v)
	g.AddEdge(0, 1)
	fmt.Println(toString(g))
	g.AddEdge(0, 0)
	fmt.Println(toString(g))
	g.AddEdge(1, 0)
	fmt.Println(toString(g))
	g.AddEdge(1, 2)
	g.AddEdge(1, 2)
	fmt.Println(toString(g))
}
