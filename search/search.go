package search

type Node struct {
	Key  interface{}
	Data interface{}
	Next *Node
	Prev *Node
}

// 无序链表
type SeqSearchST struct {
	Head *Node
	Tail *Node
	size int
}

func InitNode(key interface{}, data interface{}) *SeqSearchST {
	res := &Node{
		Key:  key,
		Data: data,
		Next: nil,
		Prev: nil,
	}

	table := &SeqSearchST{
		Head: res,
		Tail: res,
		size: 1,
	}
	return table
}

// 无序链表, 查找只能挨个遍历
func (c *SeqSearchST) Get(key interface{}) interface{} {
	node := c.Head
	for node != nil {
		if node.Key == key {
			return node.Data
		}

		node = (*node).Next
	}

	return nil
}

// 无序链表，直接将元素设置到链表结尾即可
func (c *SeqSearchST) Set(key interface{}, value interface{}) {
	node := &Node{
		Key:  key,
		Data: value,
		Next: nil,
		Prev: nil,
	}

	node.Prev = c.Tail
	c.Tail.Next = node
	c.Tail = node
	c.size++
}

// 返回无序列表size
func (c *SeqSearchST) Size() int {
	return c.size
}

// 删除某个指定的key
func (c *SeqSearchST) Del(key interface{}) {
	node := c.Head
	for node != nil {
		if node.Key == key {
			node.Prev.Next = node.Next
			return
		}

		node = node.Next
	}
}

// 不支持范围查找
