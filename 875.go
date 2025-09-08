package LeetCode

import "math"

func minEatingSpeed(piles []int, h int) int {
	check := func(eatSpeed int) bool {
		sum := 0
		for _, v := range piles {
			sum += (v + eatSpeed - 1) / eatSpeed
			if sum > h {
				return false
			}
		}
		return true
	}
	l := math.MaxInt64
	r := -1
	for _, v := range piles {
		if v < l {
			l = v
		}
		if v > r {
			r = v
		}
	}
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
