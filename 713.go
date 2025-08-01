package LeetCode

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 0 {
		return 0
	}
	count := 1
	l := 0
	res := 0
	for r := 0; r < len(nums); r++ {
		count = count * nums[r]
		for count >= k {
			count = count / nums[l]
			l++
		}
		res += r - l + 1
	}
	return res
}
