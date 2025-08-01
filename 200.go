package LeetCode

//todo 不是最优解
func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, x int, y int) {
	dir1 := []int{1, -1, 0, 0}
	dir2 := []int{0, 0, 1, -1}
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return
	}
	if grid[x][y] == '1' {
		grid[x][y] = '0'
		for i := 0; i < 4; i++ {
			dfs(grid, x+dir1[i], y+dir2[i])
		}
	}
}
