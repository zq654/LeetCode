package LeetCode

import (
	"slices"
)

func maximumBeauty(items [][]int, queries []int) []int {
	slices.SortFunc(items, func(a, b []int) int {
		return a[0] - b[0]
	})
	//空间换时间一个数组记录 [0,i]里面的最大值
	maxValueList := make([]int, len(items))
	maxValue := 0
	for i := 0; i < len(maxValueList); i++ {
		if items[i][1] > maxValue {
			maxValue = items[i][1]
		}
		maxValueList[i] = maxValue
	}
	for i := 0; i < len(queries); i++ {
		target, _ := slices.BinarySearchFunc(items, queries[i]+1, func(ints []int, i int) int {
			return ints[0] - i
		})
		if target == 0 {
			queries[i] = 0
		} else {
			queries[i] = maxValueList[target-1]
		}
	}
	return queries
}
