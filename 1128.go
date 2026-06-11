package LeetCode

func numEquivDominoPairs(dominoes [][]int) int {
	recordMap := make(map[int]int) // []int -> count
	res := 0
	for _, dominoe := range dominoes {
		dominoe[0], dominoe[1] = min(dominoe[0], dominoe[1]), max(dominoe[0], dominoe[1])
		count := dominoe[0]*10 + dominoe[1]
		res += recordMap[count]
		recordMap[count]++
	}
	return res
}
