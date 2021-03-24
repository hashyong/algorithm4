package search

const (
	Red   = true
	Black = false
)

type RBTree struct {
	Root *TreeNode
}

func isRed(node *TreeNode) bool {
	if node == nil {
		return false
	}

	return node.color == Red
}

// 就是把红色右链接变成红色左链接
// 就是 两节点中较小key做为根节点变为 较大节点做为根节点
func rotateLeft(node *TreeNode) *TreeNode {
	x := node.Right
	node.Right = x.Left
	x.Left = node
	x.color = node.color
	node.color = Red
	x.N = node.N
	node.N = 1 + size(node.Left) + size(node.Right)
	return x
}

// 就是把红色左链接变成红色右链接
// 就是 两节点中较大key做为根节点变为 较小节点做为根节点
func rotateRight(node *TreeNode) *TreeNode {
	x := node.Left
	node.Left = x.Right
	x.Right = node
	x.color = node.color
	node.color = Red
	x.N = node.N
	node.N = 1 + size(node.Left) + size(node.Right)
	return x
}

// 除了将子结点的颜色由红变黑之外，我们同时还要将父 结点的颜色由黑变红
// 为啥呢？因为树向上生长了一波， 新的生长的节点必然和之前的父节点合并了，遂是3节点， 是红链接
func flipColor(node *TreeNode) {
	node.color = Red
	node.Left.color = Black
	node.Right.color = Black
}

func (r *RBTree) Put(key BSSTI, value interface{}) {
	r.Root = rb_put(r.Root, key, value)
	r.Root.color = Black
}

// 当前版本用递归实现有问题， 不过注释描述的思想没啥问题
// 后续再更新下当前版本的删除逻辑
func (r *RBTree) DelMin() {
	// 递归删除
	// 先找到最小键
	// 直到左子树为空
	// 如果当前节点为3节点， 则直接删除即可
	// 如果当前节点为2节点，看下兄弟节点是否为3节点，是的话可以借一个键过来
	// 如果当前节点为2节点，兄弟节点为2节点, 父节点为3节点， 那就从父节点借一个过来, 合并借的节点和兄弟节点
	// 如果当前节点为2节点，兄弟节点为2节点, 父节点为2节点， 那就从父节点借一个过来, 合并借的节点和兄弟节点
	// 		此时若父节点被掏空了， 继续当前借节点逻辑, 补充当前空节点
	// 		此时若根节点空了， 则树的高度下降一层
	r.Root = delRBTMin(r.Root)
}

func delRBTMin(node *TreeNode) *TreeNode {
	// 假如当前节点的左子树的左子树为空，这当前节点的左子树为待删除节点
	if node.Left == nil {
		// 假如当前节点为3节点， 则直接删除即可
		return node.Right
	}

	node.Left = delRBTMin(node.Left)
	node.N = size(node.Left) + size(node.Right) + 1
	return node
}

func rb_put(node *TreeNode, key BSSTI, value interface{}) *TreeNode {
	if node == nil {
		return &TreeNode{
			Key:   key,
			Data:  value,
			Left:  nil,
			Right: nil,
			N:     1,
			color: Red,
		}
	}

	cmp := key.Compare(node.Key)
	// key < node.key
	if cmp < 0 {
		node.Left = put(node.Left, key, value)
	} else if cmp > 0 {
		node.Right = put(node.Right, key, value)
	} else {
		node.Data = value
	}

	if isRed(node.Right) && !isRed(node.Left) {
		node = rotateLeft(node)
	}

	if isRed(node.Left) && isRed(node.Left.Left) {
		node = rotateRight(node)
	}

	if isRed(node.Left) && isRed(node.Right) {
		flipColor(node)
	}

	node.N = size(node.Left) + size(node.Right) + 1
	return node
}

// rb del code
