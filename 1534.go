package LeetCode

func countGoodTriplets1(arr []int, a int, b int, c int) int {
	res := 0
	//暴力法
	n := len(arr)
	for j := 1; j < n-2; j++ {
		for i := 1; i < j; i++ {
			for k := n - 1; k > j; k-- {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					res++
				}
			}
		}
	}
	return res
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	res := 0
	//枚举中间值 暴力发

	return res
}
