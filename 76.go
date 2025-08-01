package LeetCode

// todo 这题是困难题 自己写了Dif优化时间到0ms 开心
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	tArr := make([]int, 64)
	dif := 0
	for i := 0; i < len(t); i++ {
		if tArr[t[i]-'A'] == 0 {
			dif++
		}
		tArr[t[i]-'A']++
	}
	l := 0
	res := ""
	for r := 0; r < len(s); r++ {
		tArr[s[r]-'A']--
		if tArr[s[r]-'A'] == 0 {
			dif--
		}
		for dif == 0 {
			if r-l+1 < len(res) || res == "" {
				res = s[l : r+1]
			}
			tArr[s[l]-'A']++
			if tArr[s[l]-'A'] == 1 {
				dif++
			}
			l++
		}
	}
	return res
}

// 下面这是第一版 时间比较慢 310ms
func minWindow_2(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	tCharArr := make(map[byte]int)
	for i, _ := range t {
		tCharArr[t[i]]++
	}
	sCharArr := make(map[byte]int)
	res := ""
	l := 0
	for r := 0; r < len(s); r++ {
		sCharArr[s[r]]++
		for checkArr(sCharArr, tCharArr) {
			if r-l+1 < len(res) || res == "" {
				res = s[l : r+1]
			}
			sCharArr[s[l]]--
			l++
		}
	}
	return res
}

func checkArr(sArr map[byte]int, tArr map[byte]int) bool {
	for k, _ := range tArr {
		if tArr[k] > sArr[k] {
			return false
		}
	}
	return true
}
