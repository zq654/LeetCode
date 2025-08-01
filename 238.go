package LeetCode

// todo err
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = 1
	}
	left := 0
	right := len(nums) - 1
	leftCount := 1
	rightCount := 1
	for left < len(nums) {
		res[left] = res[left] * leftCount
		res[right] = res[right] * rightCount
		leftCount = leftCount * nums[left]
		rightCount = rightCount * nums[right]
		left++
		right--
	}
	return res
}
