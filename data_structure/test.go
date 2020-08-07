package main

import (
	"fmt"
	"go-lc/binary_tree"
)

func main() {
	pre(generateTree2())
}

func pre(tree *binary_tree.TreeNode) {
	result := make([]int, 0)
	stack := make([]binary_tree.TreeNode, 0)
	for {
		if tree.Left !=nil{
			result = append(result, tree.Val)
			stack = append(stack, *tree)
			tree = tree.Left
		}else{
			result = append(result, tree.Val)
			if tree.Right == nil{

			}else{
				tree = tree.Right
			}
		}
	}
	fmt.Println(result)
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

func generateTree2() *binary_tree.TreeNode {
	return &binary_tree.TreeNode{
		Val: 1,
		Left: &binary_tree.TreeNode{
			Val: 2,
			Left: &binary_tree.TreeNode{
				Val: 3,
				Left: &binary_tree.TreeNode{
					Val: 5,
				},
				Right: &binary_tree.TreeNode{
					Val: 6,
				},
			},
			Right: &binary_tree.TreeNode{
				Val: 4,
			},
		},
	}
}
