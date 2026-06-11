package LeetCode

func numberOfRightTriangles(grid [][]int) int64 {
	//三角形的数量 等于枚举每一个单元格 若为一 则 可以组成 （行1的数量-1）*（列1的数量-1）个三角形
	rowCountArr := make([]int, len(grid))
	colCountArr := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				rowCountArr[i]++
				colCountArr[j]++
			}
		}
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res += (rowCountArr[i] - 1) * (colCountArr[j] - 1)
			}
		}
	}
	return int64(res)
}
