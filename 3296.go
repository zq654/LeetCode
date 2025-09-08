package LeetCode

import (
	"math"
)

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	check := func(allTime int) bool {
		count := 0
		for _, v := range workerTimes {
			count += int((math.Sqrt(float64(8*allTime/v+1)) - 1) / 2)
			if count >= mountainHeight {
				return true
			}
		}
		return false
	}
	minTime := math.MaxInt64
	for _, v := range workerTimes {
		if v < minTime {
			minTime = v
		}
	}
	l := 1
	r := mountainHeight * (mountainHeight + 1) * minTime / 2

	for l <= r {
		mid := l + (r-l)>>2
		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return int64(l)
}
