package LeetCode

func isArraySpecial(nums []int, queries [][]int) []bool {
	preArr := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		if nums[i-1]%2+nums[i]%2 == 1 {
			preArr[i] = preArr[i-1] + 1
		} else {
			preArr[i] = preArr[i-1]
		}
	}
	res := make([]bool, len(queries))
	for i, query := range queries {
		res[i] = preArr[query[1]]-preArr[query[0]] == query[1]-query[0]
	}
	return res
}
