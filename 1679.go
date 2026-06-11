package LeetCode

func maxOperations(nums []int, k int) int {
	recordMap := make(map[int]int) //value -> num
	res := 0
	for _, num := range nums {
		if recordMap[k-num] > 0 {
			res++
			recordMap[k-num]--
		} else {
			recordMap[num]++
		}
	}
	return res
}
