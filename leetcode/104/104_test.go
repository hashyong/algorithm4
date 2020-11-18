package leetcode

import (
	"fmt"
	"testing"

	. "algorithm4/leetcode"
)

//  3
//  / \
//  9  20
//  / \  /  \
//  1  2 15 7
//
// i * 2 i * 2 + 1
// i / 2 i/ 2 - 1
// 1 0 left
// 2 0 right
// 3 1 left
// 4 1 right
// 5 2 left
// 6 2 right

// 前序

func BuildTree(f []int) *TreeNode {
	nodes := make([]*TreeNode, len(f))

	root := TreeNode{Val: f[0]}
	nodes[0] = &root
	for index, value := range f {
		if index == 0 {
			continue
		}

		if value == 0 {
			continue
		}
		node := TreeNode{Val: value}
		nodes[index] = &node

		// right
		if 0 == index%2 {
			parent := index/2 - 1
			nodes[parent].Right = &node
			// left
		} else {
			parent := index / 2
			nodes[parent].Left = &node
		}
	}

	return &root
}

func Test104(t *testing.T) {
	f := []int{3, 9, 20, 1, 2, 15, 7}
	fmt.Println(f)

	root := BuildTree(f)
	fmt.Print("preorder:")
	preorder(root)
	fmt.Println()

	fmt.Print("inorder:")
	inorder(root)
	fmt.Println()

	fmt.Print("postorder:")
	postorder(root)
	fmt.Println()
}

func TestMaxDepth(t *testing.T) {
	f := []int{3, 9, 20, 0, 0, 15, 7}
	root := BuildTree(f)
	fmt.Println(maxDepth(root))
}

func TestLayerOrder(t *testing.T) {
	f := []int{3, 9, 20, 1, 2, 15, 7}
	root := BuildTree(f)
	fmt.Println(layorder(root))
}
