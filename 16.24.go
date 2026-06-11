package LeetCode

func pairSums(nums []int, target int) [][]int {
	recordMap := make(map[int]int) //value -> num
	res := make([][]int, 0)
	for _, num := range nums {
		if recordMap[target-num] > 0 {
			res = append(res, []int{num, target - num})
			recordMap[target-num]--
		} else {
			recordMap[num]++
		}
	}
	return res
}
