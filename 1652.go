package LeetCode

func decrypt(code []int, k int) []int {
	count := 0
	res := make([]int, len(code))
	if k == 0 {
		return res
	}
	//先定位窗口开始
	var begin int
	if k > 0 {
		begin = 1
	} else {
		k = -k
		begin = len(code) - k
	}
	for i := 0; i < len(code)+k; i++ {
		count += code[(i+begin)%len(code)]
		if i < k-1 {
			continue
		}
		res[(i-k+1)%len(code)] = count
		count -= code[(i+begin-k+1)%len(code)]
	}
	return res
}
