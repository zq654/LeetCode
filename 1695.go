package LeetCode

func maximumUniqueSubarray(nums []int) int {
	//count 记录窗口分数
	//map 记录某个值出现的次数
	count := 0
	res := 0
	numMap := make(map[int]int)
	l := 0
	for r := 0; r < len(nums); r++ {
		count += nums[r]
		numMap[nums[r]]++
		for numMap[nums[r]] > 1 {
			numMap[nums[l]]--
			count -= nums[l]
			l++
		}
		res = max(res, count)
	}
	return res
}
