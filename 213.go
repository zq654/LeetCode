package LeetCode

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	//直接判断掐头大还是去尾大
	return max(dp(nums[1:]), dp(nums[:len(nums)-1]))
}
func dp(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	length := len(nums)
	dpArr := make([]int, length)
	dpArr[0] = nums[0]
	dpArr[1] = max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		dpArr[i] = max(dpArr[i-2]+nums[i], dpArr[i-1])
	}
	return dpArr[length-1]
}
