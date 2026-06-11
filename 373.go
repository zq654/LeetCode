package LeetCode

import (
	"slices"
	"sort"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	//二分查找最小sum 使满足 俩数之和小于等于 sum的个数至少有k个
	check := func(sum int) bool {
		count := 0
		for i := 0; i < len(nums1); i++ {
			count += sort.SearchInts(nums2, sum-nums1[i]+1)
		}
		return count >= k
	}
	res := make([][]int, 0)
	count := 1
	sum := sort.Search(nums1[len(nums1)-1]+nums2[len(nums2)-1], check)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i]+nums2[j] <= sum {
				count++
				res = append(res, []int{nums1[i], nums2[j]})
			}
		}
	}
	slices.SortFunc(res, func(a, b []int) int {
		return a[0] + a[1] - b[0] - b[1]
	})
	return res[0:k]
}
