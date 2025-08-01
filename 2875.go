package LeetCode

func minSizeSubarray(nums []int, target int) int {
	total := 0
	for _, i := range nums {
		total += i
	}
	count := (target / total)
	residue := (target % total)
	result := append(nums, nums...)
	n := len(result)
	l, r, ans := 0, 0, 100000
	win := 0
	for r < n {
		win += result[r]
		r++
		for win > residue {
			win -= result[l]
			l++
		}
		if win == residue {
			ans = min(ans, r-l)
		}
	}
	if ans == 100000 {
		return -1
	}
	return count*(len(nums)) + ans
}
