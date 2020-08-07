package binary_tree

import (
	"fmt"
	"go-lc/binary_tree"
)

func Tree() {
	tree := generateTree()
	binary_tree.PreorderTraversal(tree)
	fmt.Println()
	for _, re := range binary_tree.PreorderTraversal2(tree) {
		fmt.Print(re)
	}
	fmt.Println()
	for _, re := range binary_tree.InorderTraversal(tree) {
		fmt.Print(re)
	}
	fmt.Println()
	for _, re := range binary_tree.PostOrderTraversal(tree) {
		fmt.Print(re)
	}
}

func generateTree() *binary_tree.TreeNode {
	return &binary_tree.TreeNode{
		Val: 2,
		Left: &binary_tree.TreeNode{
			Val:  3,
			Left: nil,
			Right: &binary_tree.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &binary_tree.TreeNode{
			Val: 5,
			Left: &binary_tree.TreeNode{
				Val: 6,
				Left: &binary_tree.TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &binary_tree.TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
}
