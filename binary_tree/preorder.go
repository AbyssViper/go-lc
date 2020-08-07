package binary_tree

import "fmt"

// 前序递归
func PreorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val)
	PreorderTraversal(root.Left)
	PreorderTraversal(root.Right)
}

// 前序非递归
func PreorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		// 一直遍历向左遍历到底，把对应的左斜到底的 root值都存储起来
		for root != nil {
			// 由于前序，当前的root是第一个需要记录的，
			result = append(result, root.Val)
			// 当前节点入栈，后续还需要使用其右分支
			stack = append(stack, root)
			root = root.Left
		}
		// 当前最深的右边节点（先出栈的）
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 往右边分支移动，下一次循环，再不断往左遍历这个右分支
		root = node.Right
	}
	return result
}
