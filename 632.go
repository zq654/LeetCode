package LeetCode

import (
	"container/heap"
	"math"
	"slices"
)

func smallestRange(nums [][]int) []int {
	numArr := make([][2]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			numArr = append(numArr, [2]int{nums[i][j], i}) //值 行数
		}
	}
	slices.SortFunc(numArr, func(a, b [2]int) int {
		return a[0] - b[0]
	})
	l := 0
	dif := 0
	res := []int{0, math.MaxInt}
	numMap := make(map[int]int, 0) //记录第n行在窗口里面有几个值
	for r := 0; r < len(numArr); r++ {
		numMap[numArr[r][1]]++
		if numMap[numArr[r][1]] == 1 {
			dif++
		}
		for dif >= len(nums) {
			if (numArr[r][0]-numArr[l][0] < res[1]-res[0]) || (numArr[r][0]-numArr[l][0] == res[1]-res[0] && numArr[r][0] < res[1]) {
				res[1] = numArr[r][0]
				res[0] = numArr[l][0]
			}
			numMap[numArr[l][1]]--
			if numMap[numArr[l][1]] == 0 {
				dif--
			}
			l++
		}
	}
	heap.Fix()
	return res
}
