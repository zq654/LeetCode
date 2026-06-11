package LeetCode

func numPairsDivisibleBy60(times []int) int {
	recordMap := make(map[int]int) //value%60 ->count
	res := 0
	for _, time := range times {
		tempValue := time % 60
		res += recordMap[(60-tempValue)%60]
		recordMap[tempValue]++
	}
	return res
}
