package base

type BagI interface {
	Add(int)
	Empty() bool
	Size() int
}

type Bag struct {
	item int
}

type QueueI interface {
	Enqueue(int)
	Deque() int
}
