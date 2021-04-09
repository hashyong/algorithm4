package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepthFirstSearch_dfs(t *testing.T) {
	g := New().GraphIn("tinyCG.json")
	fmt.Println(toString(g))

	gDFS := NewDFS()
	gDFS.DFS(g, 0)

	assert.Equal(t, gDFS.Marked(5), true)
	assert.Equal(t, gDFS.Marked(3), true)
}

func TestDepthFirstPaths_DFP(t *testing.T) {
	g := New().GraphIn("tinyCG.json")
	fmt.Println(toString(g))

	g1 := NewDFP()
	g1.DFP(g, 0)

	fmt.Println(g1.pathTo(4))
}

func TestConnectedComponent_CC(t *testing.T) {
	g := New().GraphIn("tinyCG.json")
	g1 := NewCC()
	g1.CC(g)
	fmt.Println(g1.getID(1))
	fmt.Println(g1.connected(1, 2))
	assert.Equal(t, g1.CCCount(), 1)
}
