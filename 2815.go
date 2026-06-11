package LeetCode

func maxSum(nums []int) int {
	recordMap := make(map[int]int) //最大数位 -> value
	res := -1
	for _, num := range nums {
		value := getMaxValue(num)
		if lastNum, ok := recordMap[value]; ok {
			res = max(res, lastNum+num)
		}
		if num > recordMap[value] {
			recordMap[value] = num
		}
	}
	return res
}
func getMaxValue(x int) int {
	res := -1
	for x > 0 {
		res = max(res, x%10)
		x /= 10
	}
	return res
}
