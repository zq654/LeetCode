package LeetCode

func maxValue(n int, index int, maxSum int) int {
	//index 前面有index个值 这些值最大值为 index-1 公差为-1 最小值只能是1
	//index 后面有n-1-index 个值 最大值为 index-1 公差为-1 最小值只能是1
	check := func(topValue int) bool {
		count := 0
		value0 := topValue - index         //index为0下标的值
		valueN := topValue - n + 1 + index //index为n-1的下标的值
		if value0 < 1 {
			lIndex := index - (topValue - 1)
			count += lIndex + (topValue)*(topValue-1)/2
		} else {
			count += (value0 + topValue - 1) * index / 2
		}
		if valueN < 1 {
			rIndex := index + topValue - 1
			count += n - rIndex - 1 + (topValue)*(topValue-1)/2
		} else {
			count += (valueN + topValue - 1) * (n - 1 - index) / 2
		}
		if count+topValue > maxSum {
			return false
		} else {
			return true
		}
	}

	l := 1
	r := maxSum
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}
