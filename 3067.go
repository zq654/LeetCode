package LeetCode

type edge struct {
	to     int
	weight int
}

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	//思路统计每个节点可以到哪些节点
	//因为是树因此不可能有环
	n := len(edges) + 1
	grid := make([][]edge, n)
	for _, edgeData := range edges {
		grid[edgeData[0]] = append(grid[edgeData[0]], edge{edgeData[1], edgeData[2]})
		grid[edgeData[1]] = append(grid[edgeData[1]], edge{edgeData[0], edgeData[2]})
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		edgeArr := grid[i]
		if len(edgeArr) <= 1 {
			continue
		}
		count := 0
		var dfs func(currentNode int, toNode int, weight int)
		dfs = func(currentNode int, toNode int, weight int) {
			if weight%signalSpeed == 0 {
				count++
			}
			for _, edgeData := range grid[toNode] {
				if edgeData.to != currentNode {
					dfs(toNode, edgeData.to, weight+edgeData.weight)
				}
			}
		}
		sum := 0
		for _, edgeData := range edgeArr {
			count = 0
			dfs(i, edgeData.to, edgeData.weight)
			res[i] += count * sum
			sum += count
		}
	}
	return res
}
