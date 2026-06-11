package LeetCode

func maxProfit(prices []int, strategy []int, k int) int64 {
	//想求最后利润最大,可以求被替换的那部分涨的利润最多
	//固定长度的滑动窗口
	profit := 0
	mid := k / 2
	for r := 0; r < k; r++ {
		if r < mid {
			profit -= strategy[r] * prices[r]
		} else {
			profit += (1 - strategy[r]) * prices[r]
		}
	}
	//加入一个值把他变成1 吐出一个值把他由0变成原本值
	res := max(profit, 0)
	for r := k; r < len(strategy); r++ {
		l := r - k
		profit += (1 - strategy[r]) * prices[r]
		profit += strategy[l] * prices[l]
		//l+mid 的这个值要由1变成0
		profit -= prices[l+mid]
		res = max(res, profit)
	}
	sum := 0
	for i, price := range prices {
		sum += price * strategy[i]
	}
	return int64(sum + res)
}
