package LeetCode

import "sort"

func findKthNumber(m int, n int, k int) int {
	//求k小 就是检查这个值是否满足有k个值小于等于他
	check := func(mx int) bool {
		count := 0
		for i := 1; i <= m; i++ {
			count += min(mx/m, n)
		}
		return count >= k
	}
	return sort.Search(m*n, check)
}

// 数学优化
func findKthNumber(m, n, k int) int {
	return sort.Search(m*n, func(x int) bool {
		cnt := x / n * n
		for i := x/n + 1; i <= m; i++ {
			cnt += x / i
		}
		return cnt >= k
	})
}
