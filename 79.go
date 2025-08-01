package LeetCode

//func exist(board [][]byte, word string) bool {
//
//	length := len(board)
//	width := len(board[0])
//	if length*width < len(word) {
//		return false
//	}
//	// 优化1， word里存在某个字符出现次数 大于 该字符在board中出现的次数，则不存在
//	cnt := make(map[byte]int)
//	for _, row := range board {
//		for _, col := range row {
//			cnt[col]++
//		}
//	}
//	wordCnt := make(map[byte]int)
//	w := []byte(word)
//	for _, ch := range w {
//		wordCnt[ch]++
//		if wordCnt[ch] > cnt[ch] {
//			return false
//		}
//	}
//	// 优化2，word的尾部字符在board中的出现次数 小于 word头部字符在board中的出现次数，则反转后搜索，效率更高
//	if wordCnt[w[len(w)-1]] < wordCnt[w[0]] {
//		slices.Reverse(w)
//	}
//	wordArr := []byte(word)
//
//	flag := false
//	for i := 0; i < length; i++ {
//		for j := 0; j < width; j++ {
//			if board[i][j] == wordArr[0] {
//				flag = flag || dfs(board, wordArr, 0, i, j)
//				if flag {
//					return true
//				}
//			}
//		}
//	}
//	return false
//}
//
//func dfs(board [][]byte, wordArr []byte, depth int, x int, y int) bool {
//	var direct1 = []int{1, -1, 0, 0}
//	var direct2 = []int{0, 0, 1, -1}
//	length := len(board)
//	width := len(board[0])
//	if depth == len(wordArr)-1 {
//		return true
//	}
//	flag := false
//	for dir := 0; dir < 4; dir++ {
//		oldX := x
//		oldY := y
//		oldValue := board[x][y]
//		x = x + direct1[dir]
//		y = y + direct2[dir]
//		if x >= 0 && x < length && y >= 0 && y < width && board[x][y] != '0' && board[x][y] == wordArr[depth+1] {
//			board[oldX][oldY] = '0'
//			flag = flag || dfs(board, wordArr, depth+1, x, y)
//			board[oldX][oldY] = oldValue
//		}
//		x = oldX
//		y = oldY
//	}
//	return flag
//}
