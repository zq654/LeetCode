package LeetCode

func countSubarrays(nums []int, k int) int64 {
	res := 0
	maxValue := -1
	l := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
		}
	}
	for r := 0; r < len(nums); r++ {
		if nums[r] == maxValue {
			count++
		}
		for count >= k {
			if nums[l] == maxValue {
				count--
			}
			l++
		}
		res += l
	}
	return int64(res)
}
