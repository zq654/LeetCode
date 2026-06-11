package LeetCode

// 类似冒泡
func moveZeroes(nums []int) {
	zeroIndex := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroIndex = i
			break
		}
	}
	if zeroIndex == -1 {
		return
	}
	for i := zeroIndex + 1; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[zeroIndex] = nums[zeroIndex], nums[i]
			zeroIndex++
		}
	}
}
