package LeetCode

func maxScoreSightseeingPair(values []int) int {
	//题意求values[i] + values[j] + i - j 的最大值 固定右侧 求 values[j]-j 的最大值
	res := 0
	maxValue := values[0]
	for i, value := range values {
		if i == 0 {
			continue
		}
		res = max(value+maxValue-i, res)
		maxValue = max(maxValue, value+i)
	}
	return res
}
