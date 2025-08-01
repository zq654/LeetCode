package LeetCode

func rotate(nums []int, k int) {
	k = k % len(nums)
	swap(nums, 0, len(nums)-1)
	swap(nums, 0, k-1)
	swap(nums, k, len(nums)-1)
}
func swap(nums []int, left int, right int) {
	for left < right {
		nums[left] = nums[left] ^ nums[right]
		nums[right] = nums[left] ^ nums[right]
		nums[left] = nums[left] ^ nums[right]
		left++
		right--
	}
}
