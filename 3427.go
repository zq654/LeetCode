package LeetCode

func subarraySum(nums []int) int {
	temp := make([]int, len(nums)+1)
	for i := 1; i < len(nums); i++ {
		temp[i] += temp[i-1] + nums[i-1]
	}
	sum := 0
	for i := 1; i < len(nums); i++ {
		sum += temp[i+1] - temp[max(0, i-nums[i-1])]
	}
	return sum
}
