package LeetCode

func maxNonOverlapping(nums []int, target int) int {
	//因为nums[i] 可以小于零 因此不可用滑动窗口  =》 前缀和
	//dp [i+1] 表示nums[0,i] 的数组切片里面最多有几个和为target的子数组
	dp := make([]int, len(nums)+1)
	recordMap := make(map[int]int) //期待前缀和值为X -> 期待前缀和为X的下标+1 因为可能为负数 所以贪心每次更新就可以保证这个子数组更短
	sum := 0
	recordMap[target] = 0
	for i, num := range nums {
		sum += num
		if index, ok := recordMap[sum]; ok {
			dp[i+1] = max(dp[i], dp[index]+1)
		} else {
			dp[i+1] = dp[i]
		}
		recordMap[sum+target] = i + 1
	}
	return dp[len(nums)]
}
