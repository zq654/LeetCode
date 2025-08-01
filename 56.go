package LeetCode

import (
	"slices"
)

func merge(intervals [][]int) [][]int {
	//排序
	slices.SortFunc(intervals, func(p, q []int) int {
		return p[0] - q[0]
	})
	res := make([][]int, 0)
	left := intervals[0][0]
	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		tempLeft := intervals[i][0]
		tempRight := intervals[i][1]
		if tempLeft <= right {
			right = max(tempRight, right)
		} else {
			res = append(res, []int{left, right})
			left = tempLeft
			right = tempRight
		}
	}
	res = append(res, []int{left, right})
	return res

}
