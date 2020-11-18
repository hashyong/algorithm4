package leetcode

import (
	"fmt"

	. "algorithm4/leetcode"
)

// 给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例：
// 给定二叉树 [3,9,20,null,null,15,7]，
//
//  3
//  / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
// Related Topics 树 深度优先搜索
// 👍 740 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 前序
func preorder(t *TreeNode) {
	if t == nil {
		return
	}
	fmt.Print(t.Val, "->")
	preorder(t.Left)
	preorder(t.Right)
}

// 中序
func inorder(t *TreeNode) {
	if t == nil {
		return
	}
	inorder(t.Left)
	fmt.Print(t.Val, "->")
	inorder(t.Right)
}

// 后序
func postorder(t *TreeNode) {
	if t == nil {
		return
	}
	postorder(t.Left)
	postorder(t.Right)
	fmt.Print(t.Val, "->")
}

// 层序遍历
func layorder(t *TreeNode) []int {
	var res []int
	if t == nil {
		return res
	}

	var queue []*TreeNode
	// 第一个元素入队
	queue = append(queue, t)

	for len(queue) > 0 {
		local := queue[0]
		res = append(res, local.Val)
		// 出队1个元素
		queue = queue[1:]

		if local.Left != nil {
			queue = append(queue, local.Left)
		}

		if local.Right != nil {
			queue = append(queue, local.Right)
		}
	}

	return res
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// leetcode submit region end(Prohibit modification and deletion)
