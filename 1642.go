package LeetCode

import "container/heap"

func furthestBuilding(heights []int, bricks int, ladders int) int {
	res := len(heights) - 1
	h := &myHeap{}
	heap.Init(h)
	for i := 1; i < len(heights); i++ {
		if heights[i] > heights[i-1] {
			bricks -= heights[i] - heights[i-1]
			heap.Push(h, heights[i]-heights[i-1])
		}
		if bricks < 0 && ladders > 0 {
			ladders--
			bricks += heap.Pop(h).(int)
		}
		if bricks < 0 {
			res = i - 1
			break
		}
	}
	return res
}

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v any) {
	*h = append(*h, v.(int))
}
