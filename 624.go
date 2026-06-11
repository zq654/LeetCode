package LeetCode

func maxDistance(arrays [][]int) int {
	minValue := arrays[0][0]
	maxValue := arrays[0][len(arrays[0])-1]
	res := 0
	for i := 1; i < len(arrays); i++ {
		res = max(res, abs(maxValue-arrays[i][0]), abs(minValue-arrays[i][len(arrays[i])-1]))
		maxValue = max(maxValue, arrays[i][len(arrays[i])-1])
		minValue = min(minValue, arrays[i][0])
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
