package LeetCode

func maximumSubarraySum(nums []int, k int) int64 {
	//一个map记录是否存在 一个count记录最大值
	count := 0
	repeatMap := make(map[int]int)
	dif := 0
	for i := 0; i < k; i++ {
		count = count + nums[i]
		if repeatMap[nums[i]] == 0 {
			dif++
		}
		repeatMap[nums[i]]++
	}
	res := int64(0)
	if dif == k {
		res = int64(count)
	}
	right := k
	left := 0
	for right < len(nums) {
		count = count + nums[right]
		if repeatMap[nums[right]] == 0 {
			dif++
		}
		repeatMap[nums[right]]++
		count = count - nums[left]
		if repeatMap[nums[left]] == 1 {
			dif--
		}
		repeatMap[nums[left]]--
		if dif == k {
			res = max(res, int64(count))
		}
		right++
		left++
	}
	return res
}
