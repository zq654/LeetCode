package base

import (
	"math/rand"
)

func QuickSort(nums []int) {
	quickSortDfs(nums, 0, len(nums)-1)
}

func quickSortDfs(nums []int, l int, r int) {
	if l >= r {
		return
	}
	//随机选一个值
	pivotIdx := l + rand.Intn(r-l)
	pivot := nums[pivotIdx]
	//把这个值放左边
	nums[l], nums[pivotIdx] = nums[pivotIdx], nums[l]

	li := l + 1
	ri := r
	for li <= ri {
		for li <= r && nums[li] <= pivot {
			li++
		}
		for ri >= l && nums[ri] >= pivot {
			ri--
		}
		if li > ri {
			break
		}
		nums[li], nums[ri] = nums[ri], nums[li]
	}
	nums[l], nums[ri] = nums[ri], nums[l]

	//递归
	quickSortDfs(nums, l, ri-1)
	quickSortDfs(nums, ri+1, r)
}
