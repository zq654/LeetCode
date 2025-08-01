package LeetCode

func countSubarrays(nums []int, k int64) int64 {
	if k == 1 {
		return 0
	}
	count := 0
	res := 0
	l := 0
	for r := 0; r < len(nums); r++ {
		count += nums[r]
		for int64(count*(r-l+1)) >= k {
			count -= nums[l]
			l++
		}
		res += r - l + 1
	}
	return int64(res)
}
