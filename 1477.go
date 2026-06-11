package LeetCode

import "math"

func minSumOfLengths(nums []int, target int) int {
	//因为1 <= arr[i] <= 1000 是恒定大于0的 因此前缀和是恒定单调递增的
	//Q:如何解决互不重叠
	//思路 先用map记录可行策略 然后二分查找最终答案
	subArrMap := make([][]int, 0)  //可行的子数组的 长度 -> 子数组开头的下标
	recordMap := make(map[int]int) //需要的前缀和值 -> 需要这个前缀和的下标

	sum := 0
	for i, num := range nums {
		sum += num
		if index, ok := recordMap[sum]; ok {
			subArrMap = append(subArrMap, []int{index, i})
		}
		recordMap[sum+target] = i
	}
	if len(subArrMap) < 2 {
		return -1
	}
	res := math.MaxInt64
	for i := 0; i < len(subArrMap)-1; i++ {
		for j := i + 1; j < len(subArrMap); j++ {
			if subArrMap[i][1] < subArrMap[j][0] || subArrMap[j][1] < subArrMap[i][0] {
				res = min(res, subArrMap[i][1]+subArrMap[j][1]-subArrMap[j][0]-subArrMap[i][0]+2)
			}
		}
	}
	return res
}
