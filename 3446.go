package LeetCode

import "container/heap"

func sortMatrix(grid [][]int) [][]int {
	//Q:什么样的处于一条线呢？ A: i-j的值相同
	//先暴力写
	incrQueue := &myIncQueue{}
	heap.Init(incrQueue)

	decrQueue := &myDecrQueue{}
	heap.Init(decrQueue)
	//先处理左下角三角形
	n := len(grid)
	for j := 0; j < n-1; j++ {
		for i := 0; i < n-j; i++ {
			heap.Push(decrQueue, grid[i+j][j])
		}

		for i := 0; i < n-j; i++ {
			grid[i+j][j] = heap.Pop(decrQueue).(int)
		}
	}
	//再处理右上三角形
	for i := 1; i < n-1; i++ {
		for j := 0; j < n-i; j++ {
			heap.Push(incrQueue, grid[i][j+i])
		}

		for j := 0; j < n-i; j++ {
			grid[i][j+i] = heap.Pop(incrQueue).(int)
		}
	}
	return grid
}

type myIncQueue []int

type myDecrQueue struct {
	myIncQueue
}

func (m *myDecrQueue) Less(i, j int) bool {
	return (*m).myIncQueue[i]-(*m).myIncQueue[j] < 0
}

func (m *myIncQueue) Len() int {
	return len(*m)
}

func (m *myIncQueue) Less(i, j int) bool {
	return (*m)[i]-(*m)[j] > 0
}

func (m *myIncQueue) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *myIncQueue) Push(x any) {
	(*m) = append((*m), x.(int))
}

func (m *myIncQueue) Pop() (x any) {
	x = (*m)[0]
	(*m) = (*m)[1:]
	return
}
