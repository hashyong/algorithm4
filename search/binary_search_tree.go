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

// 删除最小键
func (b *BST) DelMin() {
	b.Root = delMin(b.Root)
}

// 删除最大键
func (b *BST) DelMax() {
	b.Root = delMax(b.Root)
}

// 删除指定的key
func (b *BST) Del(key BSSTI) {
	b.Root = del(b.Root, key)
}

// 范围查找， 支持查找指定范围的key
func (b *BST) Range(lo BSSTI, hi BSSTI) []BSSTI {
	var queue []BSSTI
	range1(b.Root, &queue, lo, hi)
	return queue
}

func range1(node *TreeNode, queue *[]BSSTI, lo BSSTI, hi BSSTI) {
	if node == nil {
		return
	}

	cmpLo := lo.Compare(node.Key)
	cmpHi := hi.Compare(node.Key)
	// lo < node.key
	if cmpLo < 0 {
		range1(node.Left, queue, lo, hi)
	}

	// 处理当前节点， 左子树已经处理完成
	if cmpLo <= 0 && cmpHi >= 0 {
		*queue = append(*queue, node.Key)
	}

	if cmpHi > 0 {
		range1(node.Right, queue, lo, hi)
	}
}

func del(node *TreeNode, key BSSTI) *TreeNode {
	if node == nil {
		return nil
	}

	cmp := key.Compare(node.Key)
	// key < node.key
	if cmp < 0 {
		node.Left = del(node.Left, key)
	} else if cmp > 0 {
		node.Right = del(node.Right, key)
	} else {
		// 假如右子树为空， 则返回左子树，相当于删除当前节点
		if node.Right == nil {
			return node.Left
		}

		if node.Left == nil {
			return node.Right
		}

		t := node
		node = min(t.Right)
		node.Right = delMin(t.Right)
		node.Left = t.Left
	}

	node.N = size(node.Left) + size(node.Right) + 1
	return node
}

func delMax(node *TreeNode) *TreeNode {
	if node.Right == nil {
		return node.Left
	}

	node.Right = delMax(node.Right)
	node.N = size(node.Left) + size(node.Right) + 1
	return node
}

func delMin(node *TreeNode) *TreeNode {
	// 假如左子树为空，则当前节点为待删除元素
	// 将当前节点父节点指向当前节点右子树即可
	// 当前节点无人指向会被自动个释放
	if node.Left == nil {
		return node.Right
	}

	node.Left = delMin(node.Left)
	node.N = size(node.Left) + size(node.Right) + 1
	return node
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
