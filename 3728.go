package LeetCode

func countStableSubarrays(capacity []int) int64 {
	// 前缀和 先判断有没有需求值 然后再加上当前值
	// 因为是长度必须大于三 因此更新值的时候需要落后一条数据
	// capacity[i] 并不是严格大于0 因此需要重新调整
	length := len(capacity)
	recordMap := make(map[int]map[int]int, length) //sum -> num -> count
	sum := capacity[0]
	res := 0
	for i := 1; i < length; i++ {
		if indexMap, ok := recordMap[sum-capacity[i]]; ok {
			res += indexMap[capacity[i]]
		}
		if _, ok := recordMap[sum]; !ok {
			recordMap[sum] = map[int]int{}
		}
		//更新上一条的数据 因此可以保证长度大于等于3
		recordMap[sum][capacity[i-1]]++
		sum += capacity[i]
	}
	return int64(res)
}
