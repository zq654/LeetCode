package LeetCode

func longestWPI(hours []int) int {
	//前缀和 大于8就加一 小于8就减一
	//只要存在某个值 使得当前值减去该值大于0 那这个子序列里面就有答案
	//根据贪心 子序列越长越好 因此如果有重复的值不更新
	recordMap := make(map[int]int) //value -> index
	recordMap[0] = -1
	res := 0
	sum := 0
	for i, hour := range hours {
		if hour > 8 {
			sum++
		} else {
			sum--
		}
		if sum > 0 {
			res = i + 1
		} else if index, ok := recordMap[sum-1]; ok {
			res = max(res, i-index)
		}
		if _, ok := recordMap[sum]; !ok {
			recordMap[sum] = i
		}
	}
	return res
}
