package search

const (
	Red   = true
	Black = false
)

type RBTreeNode struct {
	Key   BSSTI
	Data  interface{}
	Left  *RBTreeNode
	Right *RBTreeNode
	N     int
	color bool
}

func isRed(node *RBTreeNode) bool {
	if node == nil {
		return false
	}

	return node.color == Red
}
