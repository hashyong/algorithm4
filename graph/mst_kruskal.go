package graph

import (
	"container/heap"
	"fmt"

	"github.com/golang-collections/collections/queue"
)

// KruskalMST 克鲁斯卡尔算法(Kruskal算法)
// Prim 更适合与解决边的绸密度更高的连通网, 因为其是从顶点出发考虑的
// 和普里姆算法恰恰相反，更适合于求边稀疏的网的最小生成树。因为其是从边的角度出发的
// 算法描述也比较简单
// 将所有边都加入优先队列， 挨个出队 假如不是连通的， 则加入mst，将其连通 假如是连通的 则忽略即可
type KruskalMST struct {
	// 最小生成树的边, EdgeInterface
	mst *queue.Queue
}

func NewKruskal() *KruskalMST {
	return &KruskalMST{}
}

func (k *KruskalMST) init(g EdgeWeightedGraphInterface) {
	// init mst
	k.mst = queue.New()

	// init 优先队列
	pq := &PriorityQueue{}
	heap.Init(pq)

	// 将所有边加入到pq中
	edges := g.Edges()
	for i := edges.Front(); i != nil; i = i.Next() {
		edge := i.Value.(EdgeInterface)
		heap.Push(pq, &Item{value: edge, priority: edge.weight()})
	}

	uf := NewUF().init(g.V())
	for pq.Len() != 0 && k.mst.Len() < g.V()-1 {
		// 从pq中得到权重最小的边和顶点， 判断是否有效， 有效的话 直接加入mst即可
		// 获取权重最小的边
		e := heap.Pop(pq).(*Item).value.(EdgeInterface)
		// 获取边的俩顶点
		v := e.either()
		w, _ := e.other(v)

		// 如果已经连通 则忽略
		if uf.Connected(v, w) {
			continue
		}

		// 合并分量
		uf.QUWeightUnion(v, w)
		// 入队
		k.mst.Enqueue(e)
	}
}

func (k *KruskalMST) edges() *queue.Queue {
	return k.mst
}

func (k *KruskalMST) weight() float32 {
	var ret float32
	for tmp := k.mst.Dequeue(); tmp != nil; tmp = k.mst.Dequeue() {
		fmt.Print(" ", tmp.(EdgeInterface))
		ret += tmp.(EdgeInterface).weight()
	}
	fmt.Println(" sum = ", ret)
	return ret
}
