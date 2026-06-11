package LeetCode

func beautifulSubarrays(nums []int) int64 {
	//最高位做抵消
	//前缀和
	recordMap := make(map[int]int, len(nums)+1)
	recordMap[0] = 1
	res := 0
	sum := 0
	for _, num := range nums {
		sum ^= num
		res += recordMap[sum]
		recordMap[sum]++
	}
	return int64(res)
}
