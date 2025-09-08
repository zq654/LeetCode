package LeetCode

import "slices"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	slices.Sort(arr2)
	count := 0
	for i := 0; i < len(arr1); i++ {
		//转化数学式得到 arr1>arr2+d 并且 arr1<arr2-d
		target1 := arr1[i] - d
		target2 := arr1[i] + d
		index1 := findTarget(arr2, target1) //这样index1左侧全是小于target1的
		index2 := findTarget(arr2, target2) //index2 右侧全是大于等于target2的
		if index2 < len(arr2) && arr2[index2] == target2 {
			continue
		}
		if index1 == index2 {
			count++
		}

	}
	return count
}
func findTarget(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
