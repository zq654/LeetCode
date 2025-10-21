package LeetCode

type node struct {
	X int
	Y int
}

// 广度优先 + 二分
func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return 0
	}
	directX := []int{0, 0, 1, -1}
	directY := []int{1, -1, 0, 0}
	//遍历一遍结果统计所有存在小偷的节点
	thiefNodeList := make([]*node, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				thiefNodeList = append(thiefNodeList, &node{X: i, Y: j})
			}
		}
	}
	check := func(mx int) bool {
		queue := make([]*node, 0)
		zeroFlag := true
		for _, thiefNode := range thiefNodeList {
			if abs(thiefNode.X-0)+abs(thiefNode.Y-0) < mx {
				zeroFlag = false
				break
			}
		}
		checkMap := make([][]bool, n)
		for i := 0; i < n; i++ {
			checkMap[i] = make([]bool, n)
		}
		if zeroFlag {
			queue = append(queue, &node{X: 0, Y: 0})
			checkMap[0][0] = true
		}

		for len(queue) > 0 {
			//取头节点
			currentNode := queue[0]
			queue = queue[1:]
			if currentNode.X == n-1 && currentNode.Y == n-1 {
				return true
			}
			//添加其余节点
			for i := 0; i < 4; i++ {
				x := currentNode.X + directX[i]
				y := currentNode.Y + directY[i]
				if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] != 1 && checkMap[x][y] == false {
					checkMap[x][y] = true
					//只有满足的节点才可以放到queue中
					flag := true
					for _, thiefNode := range thiefNodeList {
						if abs(thiefNode.X-x)+abs(thiefNode.Y-y) < mx {
							flag = false
							break
						}
					}
					if flag {
						queue = append(queue, &node{X: x, Y: y})
					}
				}
			}
		}
		return false
	}

	l := 1
	r := n
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
