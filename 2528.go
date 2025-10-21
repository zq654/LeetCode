package LeetCode

import "slices"

// 二分 + 定长滑动窗口 + 贪心 把发电站建在当前站点能够得到的右侧最远部分
func maxPower(stations []int, r int, k int) int64 {
	length := len(stations)

	check := func(mx int) bool {
		//先算第一个窗口的大小
		tempArr := make([]int, length)
		tempK := k
		copy(tempArr, stations)
		count := 0
		for i := 0; i < r; i++ {
			count += tempArr[i]
		}
		l := 0
		for i := 0; i < length; i++ {
			nextIndex := i + r
			if nextIndex < length {
				count += tempArr[nextIndex]
			}
			if i >= r+1 {
				count -= tempArr[l]
				l++
			}
			if count < mx {
				tempK -= mx - count
				nextIndex = min(length-1, i+r)
				tempArr[nextIndex] += mx - count
				count = mx
				if tempK < 0 {
					return false
				}
			}
		}
		return true
	}

	left := slices.Min(stations)
	right := slices.Max(stations)*(r+1) + k

	for left <= right {
		mid := left + (right-left)>>1
		if check(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return int64(right)
}
