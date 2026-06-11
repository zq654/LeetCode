package LeetCode

// 暴力就完事了
func countOdds(low int, high int) int {
	//每两个里面肯定有一个奇数 一个偶数
	//如果是偶数开头就可以分成 【奇数 偶数】 为一组
	//如果是奇数开头就可以分成 【偶数 奇数】 为一组
	//统计 high 到 low 之间有几组 加上 low 和 high 的数量就是结果
	return low%2 + (high-low+(low+1)%2)/2
}
