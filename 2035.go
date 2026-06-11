package LeetCode

import (
	"container/heap"
	"fmt"
)

func minimumDifference(nums []int) int {
	//直接求sum/2 然后从nums里面取length/2个值 看看哪个怎么样离sum/2 更近
	//dp[i][j]代表选到第i个值的时候已经选了j个值的sum
	//j一定小于等于i&&小于等于n
	//length-1-i代表剩余的数字中最多可以选中的数量
	//如果j+length-1-i<=length/2 则这个值必须选
	sum := 0
	for _, num := range nums {
		sum += num
	}

	length := len(nums)
	n := length / 2
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, n+1)
	}
	//dp[i][0]肯定是0
	for i := 0; i < length; i++ {
		for j := 1; j <= n && j <= i; j++ {
			if j+length-1-i > n {
				dp[i][j] = getMinValue(dp[i-1][j-1]+nums[i], dp[i-1][j], sum/2)
			} else {
				dp[i][j] = dp[i-1][j-1] + nums[i]
			}
		}
	}
	fmt.Println(dp[length-1][n])
	return abs(2 * ((sum / 2) - dp[length-1][n]))
}

func getMinValue(v1 int, v2 int, target int) int {
	if abs(v1-target) < abs(v2-target) {
		return v1
	}
	return v2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var dir4 = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func trapRainWater(heightMap [][]int) (ans int) {
	m, n := len(heightMap), len(heightMap[0])
	h := hp{}
	for i, row := range heightMap {
		for j, height := range row {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				h = append(h, cell{height, i, j})
				row[j] = -1 // 标记 (i,j) 访问过
			}
		}
	}
	heap.Init(&h)

	for len(h) > 0 {
		c := heap.Pop(&h).(cell)              // 去掉短板
		minHeight, i, j := c.height, c.x, c.y // minHeight 是木桶的短板
		for _, d := range dir4 {
			x, y := i+d.x, j+d.y                                            // (i,j) 的邻居
			if 0 <= x && x < m && 0 <= y && y < n && heightMap[x][y] >= 0 { // (x,y) 没有访问过
				// 如果 (x,y) 的高度小于 minHeight，那么接水量为 minHeight - heightMap[x][y]
				ans += max(minHeight-heightMap[x][y], 0)
				// 给木桶新增一块高为 max(minHeight, heightMap[x][y]) 的木板
				heap.Push(&h, cell{max(minHeight, heightMap[x][y]), x, y})
				heightMap[x][y] = -1 // 标记 (x,y) 访问过
			}
		}
	}
	return
}

type cell struct{ height, x, y int }
type hp []cell

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].height < h[j].height }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(cell)) }
func (h *hp) Pop() any {
	res := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return res
}
