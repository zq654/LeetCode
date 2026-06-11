package LeetCode

func maximumSum(nums []int) int {
	recordMap := make(map[int]int) //sum -> value
	res := -1
	for _, num := range nums {
		sum := getSum(num)
		if lastNum, ok := recordMap[sum]; ok {
			res = max(res, lastNum+num)
		}
		recordMap[sum] = max(recordMap[sum], num)
	}
	return res
}

func getSum(x int) int {
	res := 0
	for x > 0 {
		res += x % 10
		x /= 10
	}
	return res
}
