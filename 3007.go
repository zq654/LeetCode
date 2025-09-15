package LeetCode

func findMaximumNumber(k int64, x int) int64 {
	//暴力
	check := func(value int) bool {
		tempValue := value << 1
		count := 0
		for tempValue > 0 {
			tempValue = tempValue >> x
			count++
		}
		sumCount := 0
		for value > 0 {
			value = value >> 1
			sumCount++
		}
	}

	l := 1
	r := int(k) * x
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return int64(l - 1)
}
