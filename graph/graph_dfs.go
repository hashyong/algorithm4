package graph

import (
	"fmt"

	"github.com/golang-collections/collections/stack"

	_ "github.com/golang-collections/collections"
)

// 深度优先搜索
type DepthFirstSearch struct {
	marked []bool
	count  int
}

func NewDFS() *DepthFirstSearch {
	return &DepthFirstSearch{}
}

func (d *DepthFirstSearch) DFS(g Interface, s int) {
	d.marked = make([]bool, g.V())
	d.dfs(g, s)
}

func (d *DepthFirstSearch) dfs(g Interface, v int) {
	d.marked[v] = true
	d.count++
	lis := g.Adj(v)
	for w := lis.Front(); w != nil; w = w.Next() {
		idx := w.Value.(int)
		if !d.marked[idx] {
			d.dfs(g, idx)
		}
	}
}

func (d *DepthFirstSearch) Marked(w int) bool {
	return d.marked[w]
}

func (d *DepthFirstSearch) Count() int {
	return d.count
}

// 深度优先搜索查找图中路径
type DepthFirstPaths struct {
	marked []bool
	edgeTo []int
	s      int
}

func NewDFP() *DepthFirstPaths {
	return &DepthFirstPaths{}
}

func (d *DepthFirstPaths) DFP(g Interface, s int) {
	d.marked = make([]bool, g.V())
	d.edgeTo = make([]int, g.V())
	d.s = s
	d.dfp(g, s)
}

func (d *DepthFirstPaths) dfp(g Interface, v int) {
	d.marked[v] = true
	fmt.Println(d.edgeTo)
	lis := g.Adj(v)
	for w := lis.Front(); w != nil; w = w.Next() {
		idx := w.Value.(int)
		if !d.marked[idx] {
			d.edgeTo[idx] = v
			d.dfp(g, idx)
		}
	}
}

func (d *DepthFirstPaths) hasPathTo(v int) bool {
	return d.marked[v]
}

func (d *DepthFirstPaths) pathTo(v int) []int {
	if !d.hasPathTo(v) {
		return nil
	}

	path := stack.New()
	for x := v; x != d.s; x = d.edgeTo[x] {
		path.Push(x)
	}
	path.Push(d.s)

	var ret []int
	for {
		res := path.Pop()
		if res == nil {
			return ret
		}
		ret = append(ret, res.(int))
	}
}
