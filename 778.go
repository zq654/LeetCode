package LeetCode

import "math"

var dirX = [4]int{1, -1, 0, 0}
var dirY = [4]int{0, 0, 1, -1}

func swimInWater(heights [][]int) int {
	//check 用深度优先算法若高度差小于等于mx 且能走到终点返回true 否则返回false
	check := func(mx int) bool {
		return bfs(heights, mx)
	}
	l := heights[0][0]
	r := heights[0][0]
	for _, V1 := range heights {
		for _, V2 := range V1 {
			if V2 > r {
				r = V2
			}
		}
	}

	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

// 广度优先搜索实现
func bfs(heights [][]int, mx int) bool {
	rows := len(heights)
	if rows == 0 {
		return false
	}
	cols := len(heights[0])

	// 已访问标记数组，避免重复访问
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// 队列初始化，从起点(0,0)开始
	queue := [][]int{{0, 0}}
	visited[0][0] = true

	// 目标位置
	targetX, targetY := rows-1, cols-1

	for len(queue) > 0 {
		// 取出队列头部元素
		current := queue[0]
		queue = queue[1:] // 出队
		x, y := current[0], current[1]

		// 到达终点，返回true
		if x == targetX && y == targetY {
			return true
		}

		// 探索四个方向
		for i := 0; i < 4; i++ {
			nextX := x + dirX[i]
			nextY := y + dirY[i]

			// 检查边界和访问状态
			if nextX >= 0 && nextX < rows && nextY >= 0 && nextY < cols && !visited[nextX][nextY] {
				// 检查高度差是否符合要求
				if heights[nextX][nextY] <= mx {
					visited[nextX][nextY] = true               // 标记为已访问
					queue = append(queue, []int{nextX, nextY}) // 入队
				}
			}
		}
	}

	// 遍历完所有可达节点仍未到达终点
	return false
}
