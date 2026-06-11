package LeetCode

// 不定长滑动窗口
func minimumCardPickup1(cards []int) int {
	recordMap := make(map[int]int) //value -> count
	l := 0
	res := len(cards) + 1
	for r, card := range cards {
		recordMap[card]++
		if recordMap[card] > 1 {
			for recordMap[card] > 1 {
				recordMap[cards[l]]--
				l++
			}
			res = min(res, r-l+2)
		}
	}
	if res == len(cards)+1 {
		res = -1
	}
	return res
}

// 维护不变量
func minimumCardPickup2(cards []int) int {
	recordMap := make(map[int]int) //value -> index
	res := len(cards) + 1
	for i1, card := range cards {
		if i2, ok := recordMap[card]; ok {
			res = min(res, i1-i2+1)
		}
		recordMap[card] = i1
	}
	if res == len(cards)+1 {
		res = -1
	}
	return res
}
