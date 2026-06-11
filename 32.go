package LeetCode

func longestValidParentheses(s string) int {
	bytes := []byte(s)
	left := 0
	right := 0
	res := 0

	for i := 0; i < len(s); i++ {
		if bytes[i] == '(' {
			left++
		}
		if bytes[i] == ')' {
			right++
			if right > left {
				left = 0
				right = 0
			} else if right == left {
				res = max(res, 2*right)
			}
		}
	}
	return res
}
