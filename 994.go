package LeetCode

func orangesRotting(grid [][]int) int {
	//BFS
	goodCount := 0
	badQueue := make([][2]int, 0)
	consumerTime := 0
	//先遍历填写初始化状态
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				goodCount++
			}
			if grid[i][j] == 2 {
				badQueue = append(badQueue, [2]int{i, j})
			}
		}
	}
	dir1 := []int{0, 0, 1, -1}
	dir2 := []int{1, -1, 0, 0}
	for len(badQueue) > 0 {
		tempLength := len(badQueue)
		for i := 0; i < tempLength; i++ {
			nowNode := badQueue[0]
			for dir := 0; dir < 4; dir++ {
				if nowNode[0]+dir1[dir] > 0 && nowNode[0]+dir1[dir] < len(grid) &&
					nowNode[1]+dir2[dir] > 0 && nowNode[1]+dir2[dir] < len(grid[0]) &&
					grid[nowNode[0]+dir1[dir]][nowNode[1]+dir2[dir]] == 1 {
					grid[nowNode[0]+dir1[dir]][nowNode[1]+dir2[dir]] = 2
					goodCount--
					badQueue = append(badQueue, [2]int{nowNode[0] + dir1[dir], nowNode[1] + dir2[dir]})
				}
			}
			badQueue = badQueue[1:]
		}
		consumerTime++
	}
	if goodCount > 0 {
		return -1
	}
	return consumerTime - 1
}
