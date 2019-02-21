package _104_maximum_depth_of_binary_tree

import (
	"go.intra.xiaojukeji.com/chishui/awesomeProject/kit"
)

type TreeNode = kit.TreeNode

func maxDepth(root *TreeNode) int {
	//空数直接返回 0
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var result [][]int

func levelOrder1(root *TreeNode) int {
	result = make([][]int, 0)
	if root == nil {
		return len(result)
	}
	dfsHelper(root, 0)

	return len(result)
}

func dfsHelper(node *TreeNode, level int) {
	if node == nil {
		return
	}
	if len(result) < level+1 {
		result = append(result, make([]int, 0))
	}
	result[level] = append(result[level], node.Val)
	dfsHelper(node.Left, level+1)
	dfsHelper(node.Right, level+1)
}
