package LeetCode

import (
	"math"
	"sort"
)

func preimageSizeFZF(k int) int {
	//阶乘末尾的0 单调非减
	//求k个0 即可转化成求最大的x x满足至多有k个0
	//x 里面有几个5就会有几个0
	getZero := func(x int) int {
		count := 0
		for x > 0 {
			x = x / 5
			count += x
		}
		return count
	}
	check1 := func(x int) bool {
		return getZero(x) >= k+1
	}
	check2 := func(x int) bool {
		return getZero(x) >= k
	}
	return sort.Search(math.MaxInt64, check1) - sort.Search(math.MaxInt64, check2)
}
