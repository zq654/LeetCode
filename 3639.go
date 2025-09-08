package LeetCode

func minTime(s string, order []int, k int) int {
	check := func(activeTime int) bool {
		windows := []rune(s)
		for i := 0; i <= activeTime; i++ {
			windows[order[i]] = '*'
		}
		//滑动窗口求子数组数量 越长越合法
		starCount := 0
		l := 0
		subStrNums := 0
		for r := 0; r < len(windows); r++ {
			if windows[r] == '*' {
				starCount++
			}
			for starCount > 0 {
				if windows[l] == '*' {
					starCount--
				}
				l++
			}
			subStrNums += l
			if subStrNums >= k {
				return true
			}
		}
		return false
	}
	if len(order)*(len(order)+1)/2 < k {
		return -1
	}
	l := 0
	r := len(order) - 1
	for l <= r {
		mid := l + (r-l)>>2
		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
