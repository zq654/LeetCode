package LeetCode

// len(nums) 为1e5 因此推断时间复杂度最高为n*logn
func countBadPairs(nums []int) int64 {
	//拆解题意 寻找数对 i j 满足 j-i != nums[j] - nums[i] ==> nums[i]-i != nums[j]-j
	recordMap := make(map[int]int) //nums[i]-i => count
	res := 0
	for i, num := range nums {
		target := num - i
		res += i - recordMap[target]
		recordMap[target]++
	}
	return int64(res)
}
