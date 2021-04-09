package graph

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

// DepthFirstSearch 深度优先搜索
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
		idx, _ := w.Value.(int)
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

// DepthFirstPaths 深度优先搜索查找图中路径
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

// 13
func (d *DepthFirstPaths) dfp(g Interface, v int) {
	d.marked[v] = true
	fmt.Println(d.edgeTo)
	lis := g.Adj(v)
	for w := lis.Front(); w != nil; w = w.Next() {
		idx, _ := w.Value.(int)
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

type ConnectedComponent struct {
	marked []bool
	id     []int
	// 当前有几个连通分量
	count int
}

func NewCC() *ConnectedComponent {
	return &ConnectedComponent{}
}

func (c *ConnectedComponent) CC(g Interface) {
	c.marked = make([]bool, g.V())
	c.id = make([]int, g.V())

	for s := 0; s < g.V(); s++ {
		if !c.marked[s] {
			c.dfs(g, s)
			c.count++
		}
	}
}

func (c *ConnectedComponent) dfs(g Interface, s int) {
	c.marked[s] = true
	c.id[s] = c.count

	lis := g.Adj(s)
	for w := lis.Front(); w != nil; w = w.Next() {
		idx, _ := w.Value.(int)
		if !c.marked[idx] {
			c.dfs(g, idx)
		}
	}
}

func (c *ConnectedComponent) connected(v, w int) bool {
	return c.id[v] == c.id[w]
}

func (c *ConnectedComponent) getID(v int) int {
	return c.id[v]
}

func (c *ConnectedComponent) CCCount() int {
	return c.count
}

// Cycle 使用深度优先搜索判断一课树自环的个数
type Cycle struct {
	marked   []bool
	hasCycle bool
}

// Init sdf
func (receiver Cycle) Init(g Interface) {
	receiver.marked = make([]bool, g.V())
	receiver.hasCycle = false
	for s := 0; s < g.V(); s++ {
		if !receiver.marked[s] {
			receiver.dfs(g, s, s)
		}
	}
}

func (receiver Cycle) dfs(g Interface, v int, u int) {
	receiver.marked[v] = true

	lis := g.Adj(v)
	for w := lis.Front(); w != nil; w = w.Next() {
		idx, _ := w.Value.(int)
		if !receiver.marked[idx] {
			receiver.dfs(g, idx, v)
		} else if v != u {
			receiver.hasCycle = true
		}
	}
}

func (receiver Cycle) HasCycle() bool {
	return receiver.hasCycle
}

// TwoColor 使用深度优先搜索判断是否为二分图
type TwoColor struct {
	marked     []bool
	color      []bool
	isTwoColor bool
}

func (receiver TwoColor) Init(g Interface) {
	receiver.marked = make([]bool, g.V())
	receiver.color = make([]bool, g.V())
	receiver.isTwoColor = true

	for s := 0; s < g.V(); s++ {
		if !receiver.marked[s] {
			receiver.dfs(g, s)
		}
	}
}

func (receiver TwoColor) dfs(g Interface, v int) {
	receiver.marked[v] = true

	lis := g.Adj(v)
	for w := lis.Front(); w != nil; w = w.Next() {
		idx, _ := w.Value.(int)
		if !receiver.marked[idx] {
			receiver.color[idx] = !receiver.color[v]
			receiver.dfs(g, idx)
		} else if receiver.color[idx] == receiver.color[v] {
			receiver.isTwoColor = false
		}
	}
}

func (receiver TwoColor) IsBipartite() bool {
	return receiver.isTwoColor
}
