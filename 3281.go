package LeetCode

import "slices"

func maxPossibleScore(start []int, d int) int {
	//贪心加二分
	//思路 不管具体的值 我只看下一个区间的最右侧 能不能包含住我这个区间的值 这个区间上的值贪心是max(left,lastValue+res)
	slices.Sort(start) //升序排序 start
	length := len(start)
	check := func(mx int) bool {
		lastValue := start[0]
		for i := 1; i < length; i++ {
			if start[i]+d < lastValue+mx {
				return false
			}
			lastValue = max(start[i], lastValue+mx)
		}
		return true
	}
	l := d
	r := start[length-1] - start[0]
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
