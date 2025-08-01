package LeetCode

import "strings"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	//广度优先遍历
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := make([][]int, 0)
	rolNum := 1
	rolIndex := 0
	nextRolNum := 0
	tempArr := make([]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		if node != nil {
			tempArr = append(tempArr, node.Val)
			queue = append(queue, node.Left, node.Right)
			nextRolNum = nextRolNum + 2
			rolIndex++
		}
		if node == nil {
			rolIndex++
		}
		if rolIndex == rolNum {
			rolIndex = 0
			rolNum = nextRolNum
			nextRolNum = 0
			if len(tempArr) != 0 {
				res = append(res, tempArr)
			} else {
				break
			}
			tempArr = make([]int, 0)
		}
		queue = queue[1:]
	}
	return res
}

// todo 官解更优雅
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}
