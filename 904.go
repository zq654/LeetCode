package LeetCode

func totalFruit(fruits []int) int {
	//用map记录某个品种的水果有几棵
	fruitMap := make(map[int]int)
	res := 0
	l := 0
	for r := 0; r < len(fruits); r++ {
		fruitMap[fruits[r]]++
		for len(fruitMap) > 2 {
			fruitMap[fruits[l]]--
			if fruitMap[fruits[l]] == 0 {
				delete(fruitMap, fruits[l])
			}
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}
