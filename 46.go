package LeetCode

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	for i := 0; i < length; i++ {
		for _, v := range res {
			res = append(res)
		}
	}
}
