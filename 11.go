package LeetCode

// 移动小的是因为 假如移除大的 则剩余结果永远不肯能大于 x*length
func maxArea(height []int) int {
	res := -1
	left := 0
	right := len(height) - 1
	for left < right {
		count := (right - left) * min(height[left], height[right])
		res = max(res, count)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return res
}
