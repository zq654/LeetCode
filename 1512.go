package LeetCode

func numIdenticalPairs(nums []int) int {
	recordMap := make(map[int]int) // value -> count
	res := 0
	for _, num := range nums {
		res += recordMap[num]
		recordMap[num]++
	}
	return res
}
