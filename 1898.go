package LeetCode

func maximumRemovals(s string, p string, removable []int) int {
	check := func(k int) bool {
		//recordMap := make(map[int]bool)
		//for i := 0; i <= k; i++ {
		//	recordMap[removable[i]] = true
		//}
		//把map换成数组性能提升很大
		recordMap := make([]bool, len(s))
		for i := 0; i <= k; i++ {
			recordMap[removable[i]] = true
		}
		i := 0
		j := 0
		for i < len(s) && j < len(p) {
			if s[i] == p[j] && !recordMap[i] {
				i++
				j++
			} else {
				i++
			}
		}
		if j == len(p) {
			return true
		} else {
			return false
		}
	}
	l := 0
	r := len(removable) - 1
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
