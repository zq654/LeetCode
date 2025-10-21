package LeetCode

import "math"

func minMaxWeight(n int, edges [][]int, threshold int) int {
	// threshold 作用不大 贪心只管最短的能到0的路径就行 只要能到 后续的结点都可以通过这条路到
	recordMap := make([]int, n)
	for i := 0; i < n; i++ {
		recordMap[i] = math.MaxInt64
	}
	for i := 1; i < n; i++ {
		dfs(1, edges, recordMap)
	}
	res := math.MaxInt64
	for i := 0; i < n; i++ {
		if recordMap[i] < res {
			res = recordMap[i]
		}
	}
	return res
}

func dfs(node int, edges [][]int, recordMap []int) int {
	if node == 0 {
		return math.MaxInt64
	}
	if v := recordMap[node]; v != math.MaxInt64 {
		return v
	}
	res := -1
	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		wt := edge[2]
		if from == node {
			res = min(dfs(to, edges, recordMap), wt, res)
		}
	}
	return res
}
