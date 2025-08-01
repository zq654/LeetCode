package LeetCode

func countKConstraintSubstrings(s string, k int) int {
	count0 := 0
	count1 := 0
	l := 0
	res := 0
	for r := 0; r < len(s); r++ {
		if s[r] == '0' {
			count0++
		} else {
			count1++
		}
		for count0 > k && count1 > k {
			if s[l] == '0' {
				count0--
			} else {
				count1--
			}
			l++
		}
		res += r - l + 1
	}
	return res
}
