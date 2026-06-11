package LeetCode

func numberOfBoomerangs(points [][]int) int {
	//题意求[i,j,k] dist(i,j) == dist(i,k) 没有要求 i<j<k
	for
}

func getDist(points [][]int, i int, j int) int {
	return (points[i][0]-points[j][0])*(points[i][0]-points[j][0]) + (points[i][1]-points[j][1])*(points[i][1]-points[j][1])
}
