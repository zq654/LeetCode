package LeetCode

func containsNearbyDuplicate1(nums []int, k int) bool {
	recordMap := make(map[int]int) //value -> index
	for index1, num := range nums {
		if index2, ok := recordMap[num]; ok && abs(index1-index2) <= k {
			return true
		}
		recordMap[num] = index1
	}
	return false
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 定长滑动窗口
func containsNearbyDuplicate(nums []int, k int) bool {
	recordMap := make(map[int]bool)
	for i, num := range nums {
		if recordMap[num] {
			return true
		}
		recordMap[num] = true
		if i >= k {
			recordMap[nums[i-k]] = false
		}
	}
	return false
}
