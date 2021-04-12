package graph

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
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

// DFSS 有向图多点可达性
func (d *DepthFirstSearch) DFSS(g Interface, s []int) {
	d.marked = make([]bool, g.V())
	for i := range s {
		d.dfs(g, i)
	}
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

// Cycle 使用深度优先搜索判断树是否有环
// 适用于无向图
// 核心思想是将当前节点的父节点u向下传递
// 遍历当前节点的其他边,假如某个节点已经被标记过, 理论上只能是u, 如果不是 则说明该节点已经被遍历过,从而形成环
type Cycle struct {
	marked   []bool
	hasCycle bool
}

func NewCycle() *Cycle {
	return &Cycle{}
}

// Init sdf
// 成员函数要设置为指针类型
// 指针类型可以修改 其成员变量
func (receiver *Cycle) Init(g Interface) {
	receiver.marked = make([]bool, g.V())
	receiver.hasCycle = false
	for s := 0; s < g.V(); s++ {
		if !receiver.marked[s] {
			receiver.dfs(g, s, s)
		}
	}
}

func (receiver *Cycle) dfs(g Interface, v int, u int) {
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

func (receiver *Cycle) HasCycle() bool {
	return receiver.hasCycle
}

// TwoColor 使用深度优先搜索判断是否为二分图
type TwoColor struct {
	marked     []bool
	color      []bool
	isTwoColor bool
}

func (receiver *TwoColor) Init(g Interface) {
	receiver.marked = make([]bool, g.V())
	receiver.color = make([]bool, g.V())
	receiver.isTwoColor = true

	for s := 0; s < g.V(); s++ {
		if !receiver.marked[s] {
			receiver.dfs(g, s)
		}
	}
}

func (receiver *TwoColor) dfs(g Interface, v int) {
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

func (receiver *TwoColor) IsBipartite() bool {
	return receiver.isTwoColor
}

// DirectCycle  使用深度优先搜索判断有向图是否有环
// 核心思想是 保存当前递归栈, 假如
// 一旦我们找到了一条有向边 v → w 且 w 已经存在于栈中，就找到了一个环，因为栈表示的是一条由 w 到 v 的有向路径，而 v → w 正好补全了这个环
// 同时，如果没有找到这样的边，那就意味着这幅有向图是无环的
type DirectCycle struct {
	Marked []bool
	EdgeTo []int

	// 保存当前在栈上的顶点
	OnStack []bool
	// 环的顶点
	Cycle *stack.Stack
}

func NewDirectCycle() *DirectCycle {
	return &DirectCycle{}
}

// Init sdf
// 成员函数要设置为指针类型
// 指针类型可以修改 其成员变量
func (receiver *DirectCycle) Init(g Interface) {
	receiver.Marked = make([]bool, g.V())
	receiver.EdgeTo = make([]int, g.V())
	receiver.OnStack = make([]bool, g.V())
	receiver.Cycle = stack.New()
	for s := 0; s < g.V(); s++ {
		if !receiver.Marked[s] {
			receiver.dfs(g, s)
		}
	}
}

func (receiver *DirectCycle) HasCircle() bool {
	return receiver.Cycle.Len() != 0
}

// Circle 这个接口遍历之后 栈就空了， 打印之前先copy一份较为稳妥一点，后边优化下
func (receiver *DirectCycle) Circle() {
	tmp := receiver.Cycle.Pop()
	for tmp != nil {
		fmt.Println(tmp)
		tmp = receiver.Cycle.Pop()
	}
}

func (receiver *DirectCycle) dfs(g Interface, v int) {
	receiver.Marked[v] = true
	receiver.OnStack[v] = true

	lis := g.Adj(v)
	for w := lis.Front(); w != nil; w = w.Next() {
		// 已经有环，可以直接返回了
		if receiver.HasCircle() {
			return
		}
		idx, _ := w.Value.(int)

		// 如果还没有标记过
		if !receiver.Marked[idx] {
			receiver.EdgeTo[idx] = v
			receiver.dfs(g, idx)
		} else if receiver.OnStack[idx] {
			for x := v; x != idx; x = receiver.EdgeTo[x] {
				receiver.Cycle.Push(x)
			}
			receiver.Cycle.Push(idx)
			receiver.Cycle.Push(v)
		}
	}

	receiver.OnStack[v] = false
}

// DepthFistOrder 优先级限制下的调度问题等价于计算有向无环图中的所有顶点的拓扑顺序
type DepthFistOrder struct {
	Marked []bool

	// 所有顶点的前序遍历
	Pre *queue.Queue
	// 所有顶点的后序遍历
	Post *queue.Queue
	// 所有顶点的逆后序遍历
	ReservePost *stack.Stack
}

func NewDepthFirstOrder() *DepthFistOrder {
	return &DepthFistOrder{}
}

func (d *DepthFistOrder) Init(g Interface) {
	d.Marked = make([]bool, g.V())
	d.Pre = queue.New()
	d.Post = queue.New()
	d.ReservePost = stack.New()

	for s := 0; s < g.V(); s++ {
		if !d.Marked[s] {
			d.dfs(g, s)
		}
	}
}

func (d *DepthFistOrder) dfs(g Interface, v int) {
	d.Marked[v] = true
	d.Pre.Enqueue(v)
	lis := g.Adj(v)
	for w := lis.Front(); w != nil; w = w.Next() {
		idx, _ := w.Value.(int)
		// 如果还没有标记过
		if !d.Marked[idx] {
			d.dfs(g, idx)
		}
	}
	d.Post.Enqueue(v)
	d.ReservePost.Push(v)
}

func (d *DepthFistOrder) display() {
	fmt.Println("Pre display")
	for tmp := d.Pre.Dequeue(); tmp != nil; tmp = d.Pre.Dequeue() {
		fmt.Print(" ", tmp)
	}
	fmt.Println("\nPost display")
	for tmp := d.Post.Dequeue(); tmp != nil; tmp = d.Post.Dequeue() {
		fmt.Print(" ", tmp)
	}
	fmt.Println("\nReservePost display")
	for tmp := d.ReservePost.Pop(); tmp != nil; tmp = d.ReservePost.Pop() {
		fmt.Print(" ", tmp)
	}
	fmt.Println()

}
