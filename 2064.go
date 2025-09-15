package LeetCode

func minimizedMaximum(n int, quantities []int) int {
	length := len(quantities)
	l := 1
	r := quantities[0]
	for i := 1; i < length; i++ {
		r = max(r, quantities[i])
	}
	check := func(maxValue int) bool {
		count := 0
		for i := 0; i < length; i++ {
			count += (quantities[i] + maxValue - 1) / maxValue //向上取整
			if count > n {
				return false
			}
		}
		return true
	}
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
