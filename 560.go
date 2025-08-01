package LeetCode

// todo 6/23 没写出来
func subarraySum(nums []int, k int) int {
	//前缀和 然后看看能不能得到前缀和 加k的值存不存在
	sumMap := make(map[int]int)
	res := 0
	sumMap[nums[0]] = 1
	sumMap[0] = 1
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
		sumMap[nums[i]]++
		res = res + sumMap[nums[i]-k]
	}

	return res
}
