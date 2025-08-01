package LeetCode

func longestSubarray(nums []int) int {
	res := 0
	l := 0
	//count 统计窗口里面有几个0
	count := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			count++
			for count > 1 {
				if nums[l] == 0 {
					count--
				}
				l++
			}
		}
		res = max(res, r-l)
	}
	return res
}
