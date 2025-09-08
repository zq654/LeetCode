package LeetCode

import (
	"math"
	"slices"
)

func repairCars(ranks []int, cars int) int64 {
	check := func(time int) bool {
		count := 0
		for _, v := range ranks {
			count += int(math.Sqrt(float64(time) / float64(v)))
			if count >= cars {
				return true
			}
		}
		return false
	}

	minValue := slices.Min(ranks)
	l := 1
	r := minValue * cars * cars
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

// 2 优化 能力相同的人 在规定时间内修的车的数量是一样的
func repairCars(ranks []int, cars int) int64 {
	powerList := make([]int, 101)
	minValue := ranks[0]
	for _, power := range ranks {
		powerList[power]++
		if power < minValue {
			minValue = power
		}
	}

	check := func(time int) bool {
		count := 0
		for index := minValue; index < 101; index++ {
			if powerList[index] > 0 {
				count += int(math.Sqrt(float64(time/index))) * powerList[index]
				if count >= cars {
					return true
				}
			}
		}
		return false
	}

	l := 1
	r := minValue * cars * cars
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
