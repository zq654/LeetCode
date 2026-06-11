package LeetCode

func countTrapezoids(points [][]int) int {
	//水平与x轴即取两个y值相同的点
	recordMap := make(map[int]int) //y -> count
	sum := 0                       //统计一共有多少个水平边
	res := 0                       //统计结果
	for _, point := range points {
		count := recordMap[point[1]]
		//sum - (count * (count - 1) / 2 表示除去当前y值有多少个水平与x轴的线
		//*count 是因为每多一个点就多count个相同y值的水平线
		res = (res + (sum-(count*(count-1)/2))*count)
		sum += count
		recordMap[point[1]]++
	}
	return res % (1e9 + 7)
}
