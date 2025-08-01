package LeetCode

func setZeroes(matrix [][]int) {
	lengthX := len(matrix)
	lengthY := len(matrix[0])
	mapX := make(map[int]struct{})
	mapY := make(map[int]struct{})
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				mapX[i] = struct{}{}
				mapY[j] = struct{}{}
			}
		}
	}
	for x, _ := range mapX {
		for i := 0; i < lengthY; i++ {
			matrix[x][i] = 0
		}
	}
	for y, _ := range mapY {
		for i := 0; i < lengthX; i++ {
			matrix[i][y] = 0
		}
	}
}
