package LeetCode

import (
	"container/heap"
	"sort"
)

func find132pattern(nums []int) bool {
	//查找是否有i < j < k 满足 nums[i] < nums[k] < nums[j]
	//左侧最小，右侧维护一个有序数组 然后二分查找是否存在值
	m := &myPriorityQueue{}
	heap.Init(m)

	lMinArr := make([]int, len(nums))
	lMinArr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		lMinArr[i] = min(lMinArr[i-1], nums[i])
	}
	heap.Push(m, nums[len(nums)-1])
	for i := len(nums) - 2; i >= 0; i-- {
		minValue := lMinArr[i] - 1
		if minValue >= nums[i] {
			continue
		}
		lIndex := sort.SearchInts(*m, minValue+1)
		if lIndex == len(*m) {
			continue
		}
		rIndex := sort.SearchInts(*m, nums[i]-1)
		if rIndex == -1 {
			continue
		}
		if rIndex >= lIndex {
			return true
		}
	}
	return false

}

type myPriorityQueue []int

func (m myPriorityQueue) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m myPriorityQueue) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *myPriorityQueue) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *myPriorityQueue) Pop() any {
	x := (*m)[0]
	*m = (*m)[1:]
	return x
}

func (m myPriorityQueue) Len() int {
	return len(m)
}
