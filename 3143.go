package LeetCode

func maxPointsInsideSquare(points [][]int, s string) int {
	runeArr := []rune(s)
	recordMap := make([]int, 26) //避免重复创建map
	res := 0
	check := func(length int) bool {
		count := 0
		for i := 0; i < len(points); i++ {
			if abs(points[i][0]) <= length && abs(points[i][1]) <= length {
				if recordMap[runeArr[i]-'a'] == length {
					return false
				}
				recordMap[runeArr[i]-'a'] = length
				count++
			}
		}
		res = max(res, count)
		return true
	}
	l := 0
	r := -1
	for i := 0; i < len(points); i++ {
		if abs(points[i][0]) > r {
			r = abs(points[i][0])
		}
		if abs(points[i][1]) > r {
			r = abs(points[i][1])
		}
		if points[i][0] == 0 && points[i][1] == 0 {
			res = 1
		}
	}
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	//返回点数不是边长
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
