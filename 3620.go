package LeetCode

import (
	"math"
	"slices"
)

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	dp := make([]int, len(online))
	for i := 0; i < len(online); i++ {
		dp[i] = math.MaxInt64
	}
	slices.SortFunc(edges, func(a, b []int) int {
		return a[1] - b[1]
	})
	for i := 0; i < len(edges); i++ {
		if online[edges[i][0]] && online[edges[i][1]] {
			dp[edges[i][1]] = min(dp[edges[i][1]], edges[i][2])
		}
	}
	if dp[len(online)-1] <= int(k) {
		return dp[len(online)-1]
	}
	return -1
}
