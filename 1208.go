package LeetCode

func equalSubstring(s string, t string, maxCost int) int {
	//count统计当前花了多少预算
	count := 0
	res := 0
	l := 0
	for r := 0; r < len(s); r++ {
		count += abs(int(s[r]) - int(t[r]))

		for count > maxCost {
			count -= abs(int(s[l]) - int(t[l]))
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
