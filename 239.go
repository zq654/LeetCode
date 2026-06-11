package LeetCode

// 维护滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	//单调队列
	//思路:如果队列里面有的数比当前要加进来的数还小，就可以直接舍弃队列里面的这个数了
	//因为那个数字会先出队 而且还还比这个数字小 肯定不能成为最大值
	res := make([]int, len(nums)-k+1)
	windows := make([]int, 0) //里面记录下标
	for index, num := range nums {
		//如果滑动窗口里面靠后的值比这个值还小就可以舍弃了
		//由于每进来一个就会出去一个 所以舍弃值的时候可以不用判断会不会把窗口外的值舍弃
		for len(windows) > 0 && nums[windows[len(windows)-1]] < num {
			windows = windows[:len(windows)-1]
		}
		windows = append(windows, index)
		left := index - k + 1
		if left < 0 {
			continue
		}
		//出队
		if windows[0] < left {
			windows = windows[1:]
		}
		res[left] = nums[windows[0]]
	}
	return res
}
