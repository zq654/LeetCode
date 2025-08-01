package LeetCode

func maxConsecutiveAnswers(answerKey string, k int) int {
	res := 0
	count := 0
	l := 0
	//考虑把T变成F
	for r := 0; r < len(answerKey); r++ {
		if answerKey[r] == 'T' {
			count++
		}
		for count > k {
			if answerKey[l] == 'T' {
				count--
			}
			l++
		}
		res = max(res, r-l+1)
	}
	l = 0
	count = 0
	//考虑把F变成T
	for r := 0; r < len(answerKey); r++ {
		if answerKey[r] == 'F' {
			count++
		}
		for count > k {
			if answerKey[l] == 'F' {
				count--
			}
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

// 减枝 优化
func maxConsecutiveAnswers_2(answerKey string, k int) int {
	ans, left := 0, 0
	tcnt, fcnt := 0, 0
	for right, answer := range answerKey {
		if byte(answer) == 'T' {
			tcnt++
		} else {
			fcnt++
		}
		for tcnt > k && fcnt > k {
			if answerKey[left] == 'T' {
				tcnt--
			} else {
				fcnt--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
