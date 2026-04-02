package LeetCode

import (
	"fmt"
	"slices"
)

func minMaxWeight(n int, edges [][]int, threshold int) int {
	slices.SortFunc(edges, func(a, b []int) int {
		return a[1] - b[1]
	})
	// threshold 作用不大 贪心只管最短的能到0的路径就行 只要能到 后续的结点都可以通过这条路到

	check := func(mx int) bool {
		return bfs(n, mx, edges)
	}
	l := edges[0][2]
	r := edges[0][2]
	mixValue := l
	for i := 1; i < len(edges); i++ {
		if edges[i][2] > r {
			r = edges[i][2]
		}
		if edges[i][2] < l {
			l = edges[i][2]
		}
	}
	for l <= r {
		mid := l + (r-l)/2
		if check(mid) {
			r = r - 1
		} else {
			l = l + 1
		}
	}
	if l > mixValue {
		return -1
	}
	return l
}

func bfs(n int, mx int, edges [][]int) bool {
	count := 1
	queue := make([]int, 0)
	queue = append(queue, 0)
	recordMap := make(map[int]bool)
	recordMap[0] = true

	for len(queue) > 0 {
		fmt.Printf("%d:%v   %v \n", mx, queue, recordMap)
		node := queue[0]
		queue = queue[1:]

		for i := 0; i < len(edges); i++ {
			if edges[i][1] == node && edges[i][2] <= mx && !recordMap[edges[i][0]] {
				count++
				recordMap[edges[i][0]] = true
				queue = append(queue, edges[i][0])
			}
		}
	}
	return count == n
}
