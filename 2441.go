package LeetCode

func findMaxK(nums []int) int {
	res := -1
	recordMap := make(map[int]bool)
	for _, v := range nums {
		if recordMap[-v] {
			res = max(res, abs(v))
		} else {
			recordMap[v] = true
		}
	}
	return res
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
