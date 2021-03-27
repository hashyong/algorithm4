package graph

import (
	"container/list"
	"fmt"
)

type Interface interface {
	// 下边两个方法是全局的比较好
	// 创建一个含有V个顶点但无边的图
	Graph(int) Interface
	// 从标准输入读入一幅图
	// 数据格式
	// V(顶点数量)
	// E(边的数量)
	// 1 2
	// 3 4
	// ...
	GraphIn(interface{}) Interface
	// 顶点数
	V() int
	// 边数
	E() int
	// 添加边
	AddEdge(int, int)
	// 返回和v相邻所有顶点, 返回链表的起点
	Adj(int) list.List
}

type Graph struct {
	// 顶点数目
	Ver int
	// 边的数目
	Edge int
	// 邻接表
	// 目前是使用list，也可以使用set or st
	adj []list.List
}

func New() Interface {
	g := &Graph{}
	return g
}

func (g *Graph) V() int {
	return g.Ver
}

func (g *Graph) E() int {
	return g.Edge
}

func listSearch(v *list.List, value int) bool {
	for i := v.Front(); i != nil; i = i.Next() {
		if i.Value == value {
			return true
		}
	}
	return false
}

func (g *Graph) AddEdge(v int, w int) {
	lis := &g.adj[v]
	if !listSearch(lis, w) {
		lis.PushBack(w)
		g.Edge++

	}
	lis = &g.adj[w]
	if !listSearch(lis, v) {
		lis.PushBack(v)
	}
}

func (g *Graph) Adj(i int) list.List {
	return g.adj[i]
}

func (g *Graph) Graph(v int) Interface {
	g.Ver = v
	g.Edge = 0
	g.adj = make([]list.List, v)
	return g
}

func (g *Graph) GraphIn(in interface{}) Interface {
	return g
}

// 计算v的度数
func degree(node Interface, v int) int {
	adj := node.Adj(v)
	return adj.Len()
}

// 计算所有顶点的最大度数
func maxDegree(node Interface) int {
	max := 0
	for i := 0; i < node.V(); i++ {
		if degree(node, i) > max {
			max = degree(node, i)
		}
	}
	return max
}

// 计算所有顶点的平均度数
func avgDegree(node Interface) float32 {
	return float32(2.0 * node.E() / node.V())
}

// 计算自环的个数
func numberOfSelfLoops(node Interface) int {
	count := 0
	for v := 0; v < node.V(); v++ {
		lis := node.Adj(v)
		for i := lis.Front(); i != nil; i = i.Next() {
			if v == i.Value {
				count++
			}
		}
	}
	// 每条边都被标记过2次，所以要/2
	return count / 2
}

// 字符串表示
func toString(node Interface) string {
	s := fmt.Sprintf("%d vertices, %d edges \n", node.V(), node.E())
	for v := 0; v < node.V(); v++ {
		s1 := fmt.Sprintf("%d: ", v)
		s += s1

		lis := node.Adj(v)
		for w := lis.Front(); w != nil; w = w.Next() {
			s2 := fmt.Sprintf("%d ", w.Value)
			s += s2
		}
		s += "\n"
	}
	return s
}
