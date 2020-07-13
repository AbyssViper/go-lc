package main

import (
	"fmt"
	"go-lc/data_structure/binary_tree"
)


func main() {
	tree := generateTree()
	binary_tree.PreorderTraversal(tree)
	fmt.Println()
	res := binary_tree.InorderTraversal(tree)
	for _, re := range res {
		fmt.Print(re)
	}
}

func generateTree() *binary_tree.TreeNode {
	return &binary_tree.TreeNode{
		Val:   2,
		Left:  &binary_tree.TreeNode{
			Val:   3,
			Left:  nil,
			Right: &binary_tree.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &binary_tree.TreeNode{
			Val:   5,
			Left:  &binary_tree.TreeNode{
				Val:   6,
				Left:  &binary_tree.TreeNode{
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
