package LeetCode

import (
	"sort"
)

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	check := func(mx int) bool {
		cnt, cur := 1, price[0]
		for _, x := range price {
			if x >= cur+mx {
				cur = x
				cnt++
			}
		}
		return cnt >= k
	}
	l := 1
	r := (price[len(price)-1]-price[0])/(k-1) + 1
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}
