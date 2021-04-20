package graph

import (
	"container/heap"
	"fmt"

	"github.com/golang-collections/collections/stack"
)

// Dijkstra 迪杰斯特拉(Dijkstra)算法
// 戴克斯特拉算法（英语：Dijkstra's algorithm），又译迪杰斯特拉算法，亦可不音译而称为Dijkstra算法[6]，是由荷兰计算机科学家艾兹赫尔·戴克斯特拉
// 算法比较朴素, 基于relax 操作来做
// 何为relax
// 判断一下边e, v->w, 到w的路径是 目前存储的最短距离distTo[w] 小还是  distTo[v] + e.weight()小
// 假如是后者, 则直接更新就行 否则不做任何处理
// 直到没有有效边, 整个过程结束
// 个人感觉有点贪心的意思, 每次相当于取个最小的, 不过比较对象就俩
// 接下来咱看看代码咋个写
// 1. 将起点添加至 spt树中，将其顶点list l1假如 优先队列中
// 2. 取出权重最小的顶点 v1, 优先队列出队， 将起点-v1 加入树中， 将v1对应的顶点list加入 pq中
// 3. 直至PQ中无有效边
// 一样的， 分为及时删除和延时删除两个版本
// 和Prim算法很像， 感觉原理其实挺简单, 优先队列真是个好数据结构，太方便了
type Dijkstra struct {
	// 指向对应顶点的边
	edgeTo []EdgeInterface
	// 对应顶点距离起点的距离
	distTo []float32
	// 顶点列表, 优先队列
	pq *PriorityQueue
}

func NewDijkstra() *Dijkstra {
	return &Dijkstra{}
}

const FLOATMAX = 100000

func (d *Dijkstra) init(e EdgeWeightedGraphInterface, s int) *Dijkstra {
	// init
	d.edgeTo = make([]EdgeInterface, e.V())
	d.distTo = make([]float32, e.V())
	d.pq = &PriorityQueue{}
	heap.Init(d.pq)

	// dist default
	for i := 0; i < len(d.distTo); i++ {
		d.distTo[i] = FLOATMAX
	}
	d.distTo[s] = 0.0

	heap.Push(d.pq, &Item{value: s, priority: 0.0})
	for d.pq.Len() != 0 {
		d.relax(e, heap.Pop(d.pq).(*Item).value.(int))
	}

	return d
}

func (d *Dijkstra) relax(e EdgeWeightedGraphInterface, v int) {
	// v是 要处理的顶点
	// 接下来找到v对应的边， 比较下边的权重
	lis := e.Adj(v)
	for i := lis.Front(); i != nil; i = i.Next() {
		edge := i.Value.(EdgeInterface)
		w := edge.to()

		// 新的v-w的路径小 更新当前内存结构
		if d.distTo[w] > d.distTo[v]+edge.weight() {
			d.distTo[w] = d.distTo[v] + edge.weight()
			d.edgeTo[w] = edge

			// 如果当前pq 包含当前顶点，
			ret, item := d.pq.contain(w)
			if ret {
				d.pq.update(item, w, d.distTo[w])
			} else {
				heap.Push(d.pq, &Item{value: w, priority: d.distTo[w]})
			}
		}
	}
}

func (d *Dijkstra) DistTo(v int) float32 {
	return d.distTo[v]
}

func (d *Dijkstra) hasPathTo(v int) bool {
	return d.distTo[v] < FLOATMAX
}

func (d *Dijkstra) pathTo(v int) {
	if !d.hasPathTo(v) {
		return
	}

	path := stack.New()

	for x := d.edgeTo[v]; x != nil; x = d.edgeTo[x.from()] {
		path.Push(x)
	}

	var ret []int
	for {
		res, ok := path.Pop().(EdgeInterface)
		if !ok {
			break
		}
		ret = append(ret, res.from())
	}
	ret = append(ret, v)
	fmt.Println(ret, "weight=", d.distTo[v])
}

// DijkstraAllPairsSP 任意顶点对之间的最短距离
type DijkstraAllPairsSP struct {
	all []*Dijkstra
}

func NewDijkstraAllPairsSP() *DijkstraAllPairsSP {
	return &DijkstraAllPairsSP{}
}

func (d *DijkstraAllPairsSP) Init(g EdgeWeightedGraphInterface) *DijkstraAllPairsSP {
	d.all = make([]*Dijkstra, g.V())
	for i := 0; i < g.V(); i++ {
		d.all[i] = NewDijkstra().init(g, i)
	}
	return d
}

func (d *DijkstraAllPairsSP) path(v, w int) {
	d.all[v].pathTo(w)
}

func (d *DijkstraAllPairsSP) dist(v, w int) float32 {
	return d.all[v].distTo[w]
}

// 命题 S。按照拓扑顺序放松顶点，就能在和 E+V 成正比的时间内解决无环加权有向图的单点最
// 短路径问题。
// 证明。每条边 v → w 都只会被放松一次。当 v 被放松时，得到：distTo[w]<= distTo[v]+e.
// weight()。在算法结束前该不等式都成立，因为 distTo[v] 是不会变化的（因为是按照拓扑
// 顺序放松顶点，在 v 被放松之后算法不会再处理任何指向 v 的边）而 distTo[w] 只会变小（任
// 何放松操作都只会减小 distTo[] 中的元素的值）。因此，在所有从 s 可达的顶点都被加入到
// 树中后，最短路径的最优性条件成立，命题 Q 也就成立了。时间上限很容易得到：命题 G 告诉
// 我们拓扑排序所需的时间与 E+V 成正比，而在第二次遍历中每条边都只会被放松一次，因此算
// 法总耗时与 E+V 成正比。

// AcyclicSP 说人话就是说 distTo[w] > distTo[v]+e.weight() 时相当于找到更短的路径， 需要更新
// 那假如我按照特定顺序遍历， 保证当前  distTo[w] <= distTo[v]+e.weight() 成立即可， 那就一次遍历完事， 线性实践解决
// 1. 因为 distTo[v] 是不会变化的（因为是按照拓扑顺序放松顶点，在 v 被放松之后算法不会再处理任何指向 v 的边
// 2. 而 distTo[w] 只会变小（任何放松操作都只会减小 distTo[] 中的元素的值
// 因此，在所有从 s 可达的顶点都被加入到 树中后，最短路径的最优性条件成立, 命题 Q 也就成立了
type AcyclicSP struct {
	// 指向对应顶点的边
	edgeTo []EdgeInterface
	// 对应顶点距离起点的距离
	distTo []float32
}

func (a *AcyclicSP) init(e EdgeWeightedGraphInterface, v int) *AcyclicSP {
	// init
	a.edgeTo = make([]EdgeInterface, e.V())
	a.distTo = make([]float32, e.V())

	// dist default
	for i := 0; i < len(a.distTo); i++ {
		a.distTo[i] = FLOATMAX
	}
	a.distTo[v] = 0.0

	// 取g的拓扑排序, 遍历 然后挨个relax即可
	return a
}
