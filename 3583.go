package LeetCode

func specialTriplets(nums []int) int {
	//记录j左侧值的数量和记录j右侧值的数量
	//右侧map记录整个数组值的数量 ，当index往前走的时候 map里面的值就减一
	recordRightMap := make(map[int]int)
	recordLeftMap := make(map[int]int)
	for _, num := range nums {
		recordRightMap[num]++
	}
	res := 0
	for _, num := range nums {
		recordRightMap[num]--
		res = (res + recordRightMap[2*num]*recordLeftMap[2*num]) % (1e9 + 7)
		recordLeftMap[num]++
	}
	return res
}
