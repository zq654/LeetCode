package LeetCode

func findAnagrams(s string, p string) []int {
	//长度固定的滑动窗口,滑出答案
	res := make([]int, 0)
	windows := make(map[byte]int)
	for _, b := range p {
		windows[byte(b)]++
	}
	if len(s) < len(p) {
		return make([]int, 0)
	}
	count := len(windows)
	runtimeMap := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		runtimeMap[s[i]]++
		if runtimeMap[s[i]] == windows[s[i]] {
			count--
		}
		if runtimeMap[s[i]] == windows[s[i]]+1 {
			count++
		}
	}
	if count == 0 {
		res = append(res, 0)
	}
	for i := len(p); i < len(s); i++ {
		runtimeMap[s[i]]++
		if runtimeMap[s[i]] == windows[s[i]] {
			count--
		}
		if runtimeMap[s[i]] == windows[s[i]]+1 {
			count++
		}
		runtimeMap[s[i-len(p)]]--
		if runtimeMap[s[i-len(p)]] == windows[s[i-len(p)]] {
			count--
		}
		if runtimeMap[s[i-len(p)]] == windows[s[i-len(p)]]-1 {
			count++
		}
		if count == 0 {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}
