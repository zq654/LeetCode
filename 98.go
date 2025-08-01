package LeetCode

import "math"

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, mustBigger int, mustMinner int) bool
	dfs = func(node *TreeNode, mustBigger int, mustMinner int) bool {
		if node == nil {
			return true
		}
		if (node.Left != nil && node.Left.Val >= node.Val) || (node.Right != nil && node.Right.Val <= node.Val) {
			return false
		}
		if mustBigger != math.MinInt64 && node.Val <= mustBigger {
			return false
		}
		if mustMinner != math.MaxInt64 && node.Val >= mustMinner {
			return false
		}
		return dfs(node.Left, mustBigger, node.Val) && dfs(node.Right, node.Val, mustMinner)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}
