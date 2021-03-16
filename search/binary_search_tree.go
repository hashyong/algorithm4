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

func (b *BST) Min() BSSTI {
	return min(b.Root).Key
}

func (b *BST) Floor(key BSSTI) BSSTI {
	x := floor(b.Root, key)
	if x == nil {
		return nil
	}
	return x.Key
}

// 返回排名为k的键, 即树中正好有k个键小于它
func (b *BST) Select(k int) BSSTI {
	return treeSelect(b.Root, k)
}

// 返回指定键的排名, 注意下标从0开始
func (b *BST) Rank(key BSSTI) int {
	return rank(b.Root, key)
}

func rank(node *TreeNode, key BSSTI) int {
	if node == nil {
		return 0
	}

	cmp := key.Compare(node.Key)
	// key < node.key
	if cmp < 0 {
		return rank(node.Left, key)
	}

	if cmp > 0 {
		return 1 + size(node.Left) + rank(node.Right, key)
	}

	return size(node.Left)
}

func treeSelect(node *TreeNode, k int) BSSTI {
	if node == nil {
		return nil
	}

	t := size(node.Left)
	if t > k {
		return treeSelect(node.Left, k)
	}
	if t < k {
		// -1是因为当前节点也不符合预期， 一并减掉
		return treeSelect(node.Right, k-t-1)
	}

	return node.Key
}

func floor(node *TreeNode, key BSSTI) *TreeNode {
	if node == nil {
		return nil
	}

	cmp := key.Compare(node.Key)
	if cmp == 0 {
		return node
	}

	if cmp < 0 {
		return floor(node.Left, key)
	}

	return floor(node.Right, key)

}

func min(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node
	}

	return min(node.Left)
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
