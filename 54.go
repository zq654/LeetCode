package LeetCode

var dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, 0, m*n)
	i, j := 0, -1 // 从 (0, -1) 开始
	for di := 0; len(ans) < cap(ans); di = (di + 1) % 4 {
		for x := 0; x < n; x++ { // 走 n 步（注意 n 会减少）
			i += dirs[di][0]
			j += dirs[di][1]                // 先走一步
			ans = append(ans, matrix[i][j]) // 再加入答案
		}
		n, m = m-1, n // 减少后面的循环次数
	}
	return ans
}
