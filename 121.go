package LeetCode

import "math"

func maxProfit(prices []int) int {
	minValue := prices[0]
	res := 0
	for _, price := range prices {
		if price < minValue {
			minValue = price
		}
		res = max(res, price-minValue)
	}
	return res
}
