package LeetCode

func findMaxAverage(nums []int, k int) float64 {
	count := 0
	for i := 0; i < k; i++ {
		count = count + nums[i]
	}
	r := k
	l := 0
	res := count
	for r < len(nums) {
		count = count + nums[r] - nums[l]
		res = max(res, count)
		r++
		l++
	}

	return float64(res) / float64(k)
}
