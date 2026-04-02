package main

import (
	"fmt"
	"sort"
)

func main() {
	//[1,2,3,4,5], 6
	res := TwoNumSum_2([]int{1, 2, 3, 4, 5}, 6)
	fmt.Println(res)

}
func TwoNumSum(numbers []int, target int) [][]int {
	//1 用Map记录
	//2 排序后双指针
	//用方法二
	sort.Ints(numbers)
	l := 0
	r := len(numbers) - 1
	res := make([][]int, 0)
	for l < r {
		if numbers[l]+numbers[r] < target {
			l++
		} else if numbers[l]+numbers[r] > target {
			r--
		} else {
			res = append(res, []int{numbers[l], numbers[r]})
			l++
			r--
		}
	}
	return res
}

func TwoNumSum_2(numbers []int, target int) [][]int {
	//1 用Map记录
	//2 排序后双指针
	//用方法一
	recordMap := make(map[int]struct{})
	res := make([][]int, 0)
	for _, number := range numbers {
		if _, ok := recordMap[target-number]; ok {
			res = append(res, []int{number, target - number})
		} else {
			recordMap[number] = struct{}{}
		}
	}
	return res
}

//func minMaxWeight(n int, edges [][]int, threshold int) int {
//	slices.SortFunc(edges, func(a, b []int) int {
//		return a[1] - b[1]
//	})
//	// threshold 作用不大 贪心只管最短的能到0的路径就行 只要能到 后续的结点都可以通过这条路到
//
//	check := func(mx int) bool {
//		return bfs(n, mx, edges)
//	}
//	l := math.MaxInt64
//	r := edges[0][2]
//
//	for i := 1; i < len(edges); i++ {
//		if edges[i][2] > r {
//			r = edges[i][2]
//		}
//		if edges[i][1] == 0 && edges[i][2] < l {
//			l = edges[i][2]
//		}
//	}
//	maxValue := r
//	for l <= r {
//		mid := l + (r-l)/2
//		if check(mid) {
//			r = r - 1
//		} else {
//			l = l + 1
//		}
//	}
//	if l > maxValue {
//		return -1
//	}
//	return l
//}
//
//func bfs(n int, mx int, edges [][]int) bool {
//	count := 1
//	queue := make([]int, 0)
//	queue = append(queue, 0)
//	recordMap := make([]int, n)
//	recordMap[0] = 1
//
//	for len(queue) > 0 {
//		node := queue[0]
//		queue = queue[1:]
//		i := 0
//		for ; i < len(edges); i++ {
//			if edges[i][1] == node && edges[i][2] <= mx && recordMap[edges[i][0]] == 0 {
//				count++
//				recordMap[edges[i][0]] = 1
//				queue = append(queue, edges[i][0])
//			}
//			if edges[i][1] != node {
//				break
//			}
//		}
//		edges = edges[i:]
//	}
//	return count == n
//}
