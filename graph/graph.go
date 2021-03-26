package graph

import (
	"container/list"
	"fmt"
)

type Interface interface {
	// 创建一个含有V个顶点但无边的图
	Graph(int) Interface
	// 从标准输入读入一幅图
	GraphIn(interface{}) Interface
	// 顶点数
	V() int
	// 边数
	E() int
	// 添加边
	AddEdge(int, int)
	// 返回和v相邻所有顶点
	Adj(int) []int
}

type Graph struct {
	// 顶点数目
	V int
	// 边的数目
	E int
	// 邻接表
	adj []list.List
}

// 计算v的度数
func degree(node Interface, v int) int {
	return len(node.Adj(v))
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

//计算所有顶点的平均度数
func avgDegree(node Interface) float32 {
	return float32(2.0 * node.E() / node.V())
}

// 计算自环的个数
func numberOfSelfLoops(node Interface) int {
	count := 0
	for v := 0; v < node.V(); v++ {
		for w := range node.Adj(v) {
			if v == w {
				count++
			}
		}
	}

	// 每条边都被标记过2次，所以要/2
	return count / 2
}

// 字符串表示
func toString(node Interface) string {
	s := fmt.Sprintf("%d vertices, %d edges", node.V(), node.E())
	for v := 0; v < node.V(); v++ {
		s1 := fmt.Sprintf("%d: ", v)
		s += s1
		for w := range node.Adj(v) {
			s2 := fmt.Sprintf("%d ", w)
			s += s2
		}
		s += "\n"
	}
	return s
}
