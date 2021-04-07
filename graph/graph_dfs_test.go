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
