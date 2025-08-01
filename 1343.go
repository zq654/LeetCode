package LeetCode

func numOfSubarrays(arr []int, k int, threshold int) int {
	if threshold < 1 {
		return len(arr) - k + 1
	}
	count := 0
	for i := 0; i < k; i++ {
		count = count + arr[i]
	}
	r := k
	l := 0
	target := k * threshold
	res := 0
	if count >= target {
		res++
	}
	for r < len(arr) {
		count = count + arr[r] - arr[l]
		if count >= target {
			res++
		}
		r++
		l++
	}
	return res
}
