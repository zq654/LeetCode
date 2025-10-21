package LeetCode

func minCost(n int, edges [][]int, k int) int {
	//设置一个最大值 当有权重大于最大值时 将这个权重孤立出来
	//数组记录这个node有没有被孤立出来
	recordNode := make([]int, 0)
	check := func(mx int) bool {
		count := 0
		for i := 0; i < len(edges); i++ {
			if edges[i][2] > mx {
				count++
				recordNode[edges[i][0]] = mx
				recordNode[edges[i][1]] = mx
			}

		}
	}
}
