package LeetCode

import (
	"math"
)

func shiftDistance(s string, t string, nextCost []int, previousCost []int) int64 {
	//前缀和
	const sigma = 26
	var nxtSum, preSum [sigma + 1]int
	for i := 0; i < sigma; i++ {
		nxtSum[i+1] = nxtSum[i] + nextCost[i]
		preSum[i+1] = preSum[i] + previousCost[i]
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		jump := int(t[i]) - int(s[i])
		//若jump>0 则代表往后跳jump个
		res := math.MaxInt64
		if jump > 0 {
			res = min(res, nxtSum[t[i]-'a']-nxtSum[s[i]-'a'])
			res = min(res, preSum[s[i]-'a'+1]+preSum[sigma]-preSum[t[i]-'a'+1])
		} else {
			res = min(res, preSum[s[i]-'a'+1]-preSum[t[i]-'a'+1])
			res = min(res, nxtSum[sigma]-nxtSum[s[i]-'a']+nxtSum[t[i]-'a'])
		}
		sum += res
	}
	return int64(sum)
}
