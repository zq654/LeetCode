package LeetCode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	recordMap := make(map[*TreeNode]int)
	return dfs(root, recordMap)
}
func dfs(node *TreeNode, recordMap map[*TreeNode]int) int {
	if node == nil {
		return 0
	}
	if val, ok := recordMap[node]; ok {
		return val
	}
	if node.Left == nil && node.Right == nil {
		return node.Val
	}
	count := 0
	if node.Left != nil {
		count += dfs(node.Left.Left, recordMap) + dfs(node.Left.Right, recordMap)
	}
	if node.Right != nil {
		count += dfs(node.Right.Left, recordMap) + dfs(node.Right.Right, recordMap)
	}
	res := max(dfs(node.Left, recordMap)+dfs(node.Right, recordMap), count+node.Val)
	recordMap[node] = res
	return res
}
