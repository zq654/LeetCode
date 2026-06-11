package LeetCode

func maxAbsoluteSum(nums []int) int {
	//绝对值肯定大于等于0
	res := 0
	//最后结果有可能为正值 有可能为负值
	sum1 := 0 //假设结果为正值
	sum2 := 0 //假设结果为负值
	for _, num := range nums {
		sum1 += max(sum1, 0) + num
		sum2 += min(sum2, 0) + num
		res = max(res, sum1, -sum2)
	}
	return res
}
