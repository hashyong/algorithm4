package graph

import (
	"fmt"
	"testing"
)

func TestBreadthFirstPaths_BFS(t *testing.T) {
	g := New().GraphIn("tinyCG.json")
	fmt.Println(toString(g))

	g1 := NewBFP()
	g1.BFP(g, 0)

	fmt.Println(g1.pathTo(3))
}
