package LeetCode

func maxSubarrayLength(nums []int, k int) int {
	numsMap := make(map[int]int)
	res := 0
	l := 0
	for r := 0; r < len(nums); r++ {
		numsMap[nums[r]]++
		for numsMap[nums[r]] > k {
			numsMap[nums[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}
