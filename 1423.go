package LeetCode

func maxScore(cardPoints []int, k int) int {
	leftCountArr := make([]int, k+1)
	rightCountArr := make([]int, k+1)
	for i := 1; i <= k; i++ {
		leftCountArr[i] = leftCountArr[i-1] + cardPoints[i-1]
		rightCountArr[i] = rightCountArr[i-1] + cardPoints[len(cardPoints)-i]
	}
	res := 0
	for i := 0; i <= k; i++ {
		res = max(res, leftCountArr[i]+rightCountArr[k-i], leftCountArr[i]+rightCountArr[k-i])
	}
	return res
}
