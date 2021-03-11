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
