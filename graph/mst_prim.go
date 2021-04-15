package graph

import (
	"container/heap"
	"fmt"

	"github.com/golang-collections/collections/queue"
)

// An Item is something we manage in a priority queue.
// 没办法 自己写个优先队列 抄的
// 官方文档copy的 https://golang.org/pkg/container/heap/
type Item struct {
	value    interface{} // The value of the item; arbitrary.
	priority float32     // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority float32) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// LazyPrimMST 普里姆算法（Prim算法）
type LazyPrimMST struct {
	// 最小生成树的顶点
	marked []bool
	// 最小生成树的边, EdgeInterface
	mst *queue.Queue
	// 横切边 包括失效的边, 优先队列
	pq *PriorityQueue
}

func NewLazyPrimMSt() *LazyPrimMST {
	return &LazyPrimMST{}
}

func (l *LazyPrimMST) LazyMSTInit(g EdgeWeightedGraphInterface) *LazyPrimMST {
	l.pq = &PriorityQueue{}
	l.marked = make([]bool, g.V())
	l.mst = queue.New()

	heap.Init(l.pq)

	l.visit(g, 0)

	for l.pq.Len() != 0 {
		// 获取权重最小的边
		e := heap.Pop(l.pq).(*Item).value.(EdgeInterface)

		// 获取边的俩顶点
		v := e.either()
		w, _ := e.other(v)

		// 跳过失效的边
		if l.marked[v] && l.marked[w] {
			continue
		}

		// 将边添加到树中
		l.mst.Enqueue(e)

		// 将顶点添加到树中
		if !l.marked[v] {
			l.visit(g, v)
		}
		if !l.marked[w] {
			l.visit(g, w)
		}
	}

	return l
}

func (l *LazyPrimMST) visit(g EdgeWeightedGraphInterface, v int) {
	// 标记顶点v 并将所有和v连接 未被标记的顶点的边加入pq
	// 朴素的一批
	l.marked[v] = true
	lis := g.Adj(v)
	for i := lis.Front(); i != nil; i = i.Next() {
		edge := i.Value.(EdgeInterface)
		idx, _ := edge.other(v)
		if !l.marked[idx] {
			heap.Push(l.pq, &Item{value: edge, priority: edge.weight()})
		}
	}
}

func (l *LazyPrimMST) edges() *queue.Queue {
	return l.mst
}

func (l *LazyPrimMST) weight() float32 {
	var ret float32
	for tmp := l.mst.Dequeue(); tmp != nil; tmp = l.mst.Dequeue() {
		fmt.Print(" ", tmp.(EdgeInterface))
		ret += tmp.(EdgeInterface).weight()
	}
	fmt.Println(" sum = ", ret)
	return ret
}
