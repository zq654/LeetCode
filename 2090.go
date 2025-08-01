package LeetCode

func getAverages(nums []int, k int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}
	if len(nums) < 2*k+1 {
		return res
	}
	count := 0
	for i := 0; i < 2*k+1; i++ {
		count = count + nums[i]
	}
	for i := k; i+k < len(nums); i++ {
		res[i] = count / (2*k + 1)
		if i+k+1 >= len(nums) {
			continue
		}
		count = count + nums[i+k+1]
		count = count - nums[i-k]
	}
	return res
}
