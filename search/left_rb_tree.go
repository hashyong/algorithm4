package search

const (
	Red   = true
	Black = false
)

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
