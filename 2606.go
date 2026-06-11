package LeetCode

// 贪心
func maximumCostSubstring(s string, chars string, vals []int) int {
	customerArr := make([]int, 26)
	for i := range customerArr {
		customerArr[i] = 1 + i
	}
	for i := range chars {
		customerArr[chars[i]-'a'] = vals[i]
	}
	res := 0
	sum := 0
	for i := 0; i < len(s); i++ {
		nowValue := customerArr[s[i]-'a']
		sum = max(nowValue, sum+nowValue)
		res = max(res, sum)
	}
	return res
}

// 前缀和
func maximumCostSubstring(s string, chars string, vals []int) int {
	customerArr := make([]int, 26)
	for i := range customerArr {
		customerArr[i] = 1 + i
	}
	for i := range chars {
		customerArr[chars[i]-'a'] = vals[i]
	}
	res := 0
	sum := 0
	minValue := 0
	for i := 0; i < len(s); i++ {
		sum += customerArr[s[i]-'a']
		res = max(res, sum-minValue)
		minValue = min(minValue, sum)
	}
	return res
}
