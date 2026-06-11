package LeetCode

import "math"

// 空间复杂度 O(m+n)
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

//空间复杂度 O(1)

func setZeroes(matrix [][]int) {
	//如果遇到0就把他替换拆开成 [i][0] [j][0]
	//但是matrix[0][0] 没办法知道是行为0还是列为0
	rowZero := false //标记matrix[0][0]是否行为0
	colZero := false //标记matrix[0][0]是否列为0

	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
				if i == 0 {
					colZero = true
				}
				if j == 0 {
					rowZero = true
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if rowZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	if colZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	return
}
