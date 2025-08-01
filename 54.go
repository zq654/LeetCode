package LeetCode

// todo 到时候用缩圈法写一遍
func spiralOrder(matrix [][]int) []int {
	if len(matrix[0]) == 1 {
		res := make([]int, len(matrix))
		for i, v := range matrix {
			res[i] = v[0]
		}
		return res
	}
	//定义方向
	var direct1 = []int{0, 1, 0, -1}
	var direct2 = []int{1, 0, -1, 0}
	//四个变量定义边界
	upVerge := 0
	downVerge := len(matrix) - 1
	leftVerge := 0
	rightVerge := len(matrix[0]) - 1

	dir := 0
	nums := len(matrix) * len(matrix[0])
	res := make([]int, nums)
	//坐标
	x := 0
	y := 1
	res[0] = matrix[0][0]
	for i := 1; i < nums; i++ {
		res[i] = matrix[x][y]
		if dir == 0 && x == upVerge && y == rightVerge {
			dir = (dir + 1) % 4
			upVerge++
		} else if dir == 1 && x == downVerge && y == rightVerge {
			dir = (dir + 1) % 4
			rightVerge--
		} else if dir == 2 && x == downVerge && y == leftVerge {
			dir = (dir + 1) % 4
			downVerge--
		} else if dir == 3 && x == upVerge && y == leftVerge {
			dir = (dir + 1) % 4
			leftVerge++
		}
		x = x + direct1[dir]
		y = y + direct2[dir]
	}
	return res
}
