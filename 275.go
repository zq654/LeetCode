package LeetCode

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	// 至少n个文案被引用n次
	//找到有x个文案引用次数大于等于 n 若x>=n 则返回true 若 x<n 返回false
	check := func(n int) bool {
		index := sort.SearchInts(citations, n+1) - 1
		fmt.Print("index")
		if len(citations)-index >= n {
			return true
		} else {
			return false
		}
	}
	l := 0
	r := citations[len(citations)-1]
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l + 1
}
