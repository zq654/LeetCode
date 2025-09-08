package LeetCode

import (
	"sort"
)

func maximumCandies(candies []int, k int64) int {
	//糖果可以分成的范围为 [1,min(candies)]
	//check(t)二分查找 >=t的数量是不是>=k
	check := func(candieNum int) bool {
		count := 0
		for _, v := range candies {
			count += v / candieNum
			if count >= int(k) {
				return true
			}
		}
		return false
	}
	sum := 0
	for _, v := range candies {
		sum += v
	}
	l := 1
	r := sum / int(k)

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

func maximumCandies_2(candies []int, k int64) int {
	mx, sum := 0, 0
	for _, c := range candies {
		mx = max(mx, c)
		sum += c
	}
	// 二分最大的不满足要求的 low+1，那么答案就是 low
	return sort.Search(min(mx, sum/int(k)), func(low int) bool {
		low++
		sum := 0
		for _, candy := range candies {
			sum += candy / low
		}
		return sum < int(k)
	})
}
