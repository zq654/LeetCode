package LeetCode

import "fmt"

func countPartitions(nums []int, k int) int {
	//双指针 模拟
	minValue := nums[0]
	maxValue := nums[0]
	l := 0
	r := 1
	res := 1
	for ; r < len(nums); r++ {
		num := nums[r]
		minValue = min(minValue, num)
		maxValue = max(maxValue, num)
		if maxValue-minValue > k {
			res *= r - l
			l = r
			minValue = num
			maxValue = num
		}
	}
	fmt.Println(r)
	fmt.Println(l)
	res *= r - l + 1
	return res
}
