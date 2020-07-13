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
