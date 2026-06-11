package LeetCode

// nums[i] 可以为负数因此不能用滑动窗口
// 数据范围是10的七次方 因此推理时间复杂度是nlogn
func subarraySum(nums []int, k int) int {
	//前缀和
	sum := 0
	recordMap := make(map[int]int)
	recordMap[0] = 1
	res := 0
	for _, num := range nums {
		sum += num
		res += recordMap[sum-k]
		recordMap[sum]++
	}
	return res
}
