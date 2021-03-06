package graph

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// 最小生成树的相关代码。挺有意思 学习下
// 使用边来表示图
// 每个邻接表的结点都是一个指向 Edge 对象的引用，它们含有一些冗余的信息（v 的邻接链表中的
// 每个结点都会用一个变量保存 v）。使用对象也会带来一些开销。虽然每条边的 Edge 对象都只有
// 一个，但邻接表中还是会含有两个指向同一 Edge 对象的引用。另一种广泛使用的方案是与 Graph
// 一样，用两个结点对象来表示一条边，每个结点对象都会保存顶点的信息和边的权重。这种方法也
// 是有代价的——需要两个结点，每条边的权重都会被保存两遍。

// EdgeInterface 边通用API
type EdgeInterface interface {
	// 初始化
	init(int, int, float32)
	// 边的权重
	weight() float32
	// 边两端的顶点之一
	either() int
	// 边两端的顶点之一
	other(int) (int, error)
	// 边的比较
	compareTo(edgeInterface EdgeInterface) int
	// 可视化输出
	toString() string

	// 兼容下有向图
	from() int
	to() int
}

// EdgeWeightedGraphInterface 加权无向图的API
// 也可以表示加权有向图
type EdgeWeightedGraphInterface interface {
	// Default  init
	Default(int) EdgeWeightedGraphInterface
	// In  init
	In(interface{}) EdgeWeightedGraphInterface
	// V 图的顶点数
	V() int
	// E 图的边数
	E() int
	// AddEdge 添加边
	AddEdge(edgeInterface EdgeInterface)
	// Adj 关联的所有边
	Adj(int) list.List
	// Edges 图的所有边
	Edges() list.List
	// ToString 可视化输出
	ToString() string
}

// Edge 带权重的边
type Edge struct {
	// 一个顶点
	V int
	// 另一个顶点
	W int
	// 边的权重
	Weight float32
}

func (e *Edge) init(v1, v2 int, w float32) {
	e.V = v1
	e.W = v2
	e.Weight = w
}

func (e *Edge) weight() float32 {
	return e.Weight
}

func (e *Edge) from() int {
	return e.V
}

func (e *Edge) to() int {
	return e.W
}

func (e *Edge) either() int {
	return e.V
}

func (e *Edge) other(i int) (int, error) {
	if i == e.V {
		return e.W, nil
	}

	if i == e.W {
		return e.V, nil
	}

	return 0, errors.New("not match")
}

func (e *Edge) compareTo(that EdgeInterface) int {
	ret := e.weight() - that.weight()
	switch {
	case ret < 0:
		return -1
	case ret > 0:
		return 1
	default:
		return 0
	}
}

func (e *Edge) toString() string {
	return fmt.Sprintf("[%d-%d %.2f]", e.V, e.W, e.Weight)
}

// EdgeWeightedGraph 加权无向图
// 当让 也可以表示加权有向图, 只是边不同罢了
type EdgeWeightedGraph struct {
	// 顶点的总数
	v int
	// 边的总数
	e int
	// 邻接表
	adj []list.List
	// 是否为 有向图
	// 有向图添加只添加一次即可
	isDirected bool
}

func NewEWG() *EdgeWeightedGraph {
	return &EdgeWeightedGraph{}
}

func NewDirectedEWG() *EdgeWeightedGraph {
	return &EdgeWeightedGraph{
		isDirected: true,
	}
}

func (e *EdgeWeightedGraph) Default(v int) EdgeWeightedGraphInterface {
	e.v = v
	e.e = 0
	e.adj = make([]list.List, v)
	return e
}

func (e *EdgeWeightedGraph) In(in interface{}) EdgeWeightedGraphInterface {
	type Autogenerated struct {
		V    int `json:"v"`
		Edge []struct {
			S int     `json:"s"`
			E int     `json:"e"`
			W float32 `json:"w"`
		} `json:"edge"`
	}

	// 读取标准输入数据
	content, err := ioutil.ReadFile(in.(string))
	if err != nil {
		return e
	}

	// json 转换
	var d Autogenerated
	err = json.Unmarshal(content, &d)
	if err != nil {
		return e
	}

	// 数据读取
	e.v = d.V
	e.e = len(d.Edge)
	e.adj = make([]list.List, d.V)
	for _, data := range d.Edge {
		e.AddEdge(&Edge{
			V:      data.S,
			W:      data.E,
			Weight: data.W})
	}

	return e
}

func (e *EdgeWeightedGraph) V() int {
	return e.v
}

func (e *EdgeWeightedGraph) E() int {
	return e.e
}

func (e *EdgeWeightedGraph) AddEdge(edgeInterface EdgeInterface) {
	v := edgeInterface.either()
	e.adj[v].PushBack(edgeInterface)

	// 如果是无向图 边需要添加两次
	if !e.isDirected {
		w, _ := edgeInterface.other(v)
		e.adj[w].PushBack(edgeInterface)
	}

	e.e++
}

func (e *EdgeWeightedGraph) Adj(v int) list.List {
	return e.adj[v]
}

// Edges 实现逻辑再议
// 根据某个顶点出发 做下DFS即可
func (e *EdgeWeightedGraph) Edges() list.List {
	return e.adj[0]
}

func (e *EdgeWeightedGraph) ToString() string {
	s := fmt.Sprintf("%d vertices, %d edges \n", e.V(), e.E())
	for v := 0; v < e.V(); v++ {
		s1 := fmt.Sprintf("%d: ", v)
		s += s1

		lis := e.Adj(v)
		for w := lis.Front(); w != nil; w = w.Next() {
			s2 := w.Value.(EdgeInterface).toString()
			s += s2
		}
		s += "\n"
	}
	return s
}
