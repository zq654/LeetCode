package LeetCode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	//前缀和 加上递归
	initSumMap := make(map[int]int, 0)
	initSumMap[0] = 1
	count := 0
	//先从下往上获取前缀和
	var sumFunc func(node *TreeNode, parentValue int, sumMap map[int]int)
	sumFunc = func(node *TreeNode, parentValue int, sumMap map[int]int) {
		if node == nil {
			return
		}
		node.Val += parentValue
		count += sumMap[node.Val-targetSum]
		sumMap[node.Val] = sumMap[node.Val] + 1
		sumFunc(node.Left, node.Val, sumMap)
		sumFunc(node.Right, node.Val, sumMap)
		sumMap[node.Val] = sumMap[node.Val] - 1
	}

	sumFunc(root, 0, initSumMap)
	return count
}
