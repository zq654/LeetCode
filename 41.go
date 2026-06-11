package LeetCode

func firstMissingPositive(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] <= 0 || nums[i] > len(nums) || nums[i] == i+1 || nums[nums[i]-1] == nums[i] {
			i++
			continue
		}
		value := nums[i]
		nums[value-1], nums[i] = nums[i], nums[value-1]
	}
	for j := 0; j < len(nums); j++ {
		if nums[j] != j+1 {
			return j + 1
		}
	}
	return len(nums) + 1
}
