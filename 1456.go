package LeetCode

func maxVowels(s string, k int) int {
	arr := []byte(s)
	count := 0
	for i := 0; i < k; i++ {
		if arr[i] == 'a' || arr[i] == 'e' || arr[i] == 'i' || arr[i] == 'o' || arr[i] == 'u' {
			count++
		}
	}
	res := count
	r := k
	l := 0
	for r < len(s) {
		if arr[r] == 'a' || arr[r] == 'e' || arr[r] == 'i' || arr[r] == 'o' || arr[r] == 'u' {
			count++
		}
		if arr[l] == 'a' || arr[l] == 'e' || arr[l] == 'i' || arr[l] == 'o' || arr[l] == 'u' {
			count--
		}
		r++
		l++
		res = max(res, count)
	}
	return res
}
