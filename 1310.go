package LeetCode

func xorQueries(arr []int, queries [][]int) []int {
	sum := make([]int, len(arr)+1)
	for i, w := range arr {
		sum[i+1] = sum[i] ^ w
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = sum[q[1]+1] ^ sum[q[0]]
	}
	return ans

}
