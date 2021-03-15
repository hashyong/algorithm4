package search

type TreeNode struct {
	Key   BSSTI
	Data  interface{}
	Left  *TreeNode
	Right *TreeNode
	N     int
}

type BST struct {
	Root *TreeNode
}

func InitBST() *BST {
	return &BST{
		Root: nil,
	}
}

func (b *BST) Size() int {
	return size(b.Root)
}

func size(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.N
}

func (b *BST) Get(key BSSTI) interface{} {
	return get(b.Root, key)
}

func (b *BST) Put(key BSSTI, value interface{}) {
	b.Root = put(b.Root, key, value)
}

func put(node *TreeNode, key BSSTI, value interface{}) *TreeNode {
	if node == nil {
		return &TreeNode{
			Key:   key,
			Data:  value,
			N:     1,
			Left:  nil,
			Right: nil,
		}
	}

	cmp := key.Compare(node.Key)
	if cmp > 0 {
		node.Right = put(node.Right, key, value)
	} else if cmp < 0 {
		node.Left = put(node.Left, key, value)
	} else {
		node.Data = value
	}

	node.N = size(node.Left) + size(node.Right) + 1
	return node
}

func get(x *TreeNode, key BSSTI) interface{} {
	if x == nil {
		return nil
	}

	cmp := key.Compare(x.Key)
	if cmp > 0 {
		return get(x.Right, key)
	}
	if cmp < 0 {
		return get(x.Left, key)
	}
	return x.Data
}
