package LeetCode

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	//这题转化题意 可以移动k次就代表可以进行k次数组合并
	//也就是可以转化成定长滑动窗口就最大值 窗口大小为k+1
	//第一步统计不同子区间大小
	length := len(startTime)
	timeArr := make([]int, length+1)
	end := 0
	for i := 0; i < length; i++ {
		timeArr[i] = startTime[i] - end
		end = endTime[i]
	}
	//最后面还有一个区间
	timeArr[length] = eventTime - endTime[length-1]
	res := 0
	sum := 0
	for i, num := range timeArr {
		sum += num
		if i < k {
			continue
		}
		res = max(res, sum)
		sum -= timeArr[i-k]
	}
	return res
}
