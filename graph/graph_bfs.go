package graph

import (
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

type BreadthFirstPaths struct {
	marked []bool
	edgeTo []int
	start  int
}

func NewBFP() *BreadthFirstPaths {
	return &BreadthFirstPaths{}
}

func (b *BreadthFirstPaths) BFP(g Interface, s int) {
	b.marked = make([]bool, g.V())
	b.edgeTo = make([]int, g.V())
	b.start = s
	b.bfp(g, s)
}

func (b *BreadthFirstPaths) bfp(g Interface, s int) {
	q := queue.New()
	b.marked[s] = true
	q.Enqueue(s)

	for q.Len() != 0 {
		v, _ := q.Dequeue().(int)
		lis := g.Adj(v)
		for w := lis.Front(); w != nil; w = w.Next() {
			idx, _ := w.Value.(int)
			if !b.marked[idx] {
				b.edgeTo[idx] = v
				b.marked[idx] = true
				q.Enqueue(idx)
			}
		}
	}
}

func (b *BreadthFirstPaths) hasPathTo(v int) bool {
	return b.marked[v]
}

func (b *BreadthFirstPaths) pathTo(v int) []int {
	if !b.hasPathTo(v) {
		return nil
	}

	path := stack.New()
	for x := v; x != b.start; x = b.edgeTo[x] {
		path.Push(x)
	}
	path.Push(b.start)

	var ret []int
	for {
		res := path.Pop()
		if res == nil {
			return ret
		}
		ret = append(ret, res.(int))
	}
}
