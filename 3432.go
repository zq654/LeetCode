package LeetCode

func countPartitions(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return 0
	} else {
		return len(nums) - 1
	}
}
