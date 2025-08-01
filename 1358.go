package LeetCode

// todo 模板题 没写出来
func numberOfSubstrings(s string) int {
	byteArr := make([]int, 3)
	dif := 0
	l := 0
	res := 0
	for r := 0; r < len(s); r++ {
		byteArr[s[r]-'a']++
		if byteArr[s[r]-'a'] == 1 {
			dif++
		}
		for dif == 3 {
			byteArr[s[l]-'a']--
			if byteArr[s[l]-'a'] == 0 {
				dif--
			}
			l++
		}
		res += l
	}
	return res
}
