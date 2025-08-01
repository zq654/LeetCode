package LeetCode

func minimumRecolors(blocks string, k int) int {
	count := 0
	res := 101
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			count++
		}
		if i < k-1 {
			continue
		}
		res = min(res, count)
		if blocks[i-k+1] == 'W' {
			count--
		}
	}
	return res
}
