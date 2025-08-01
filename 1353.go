package LeetCode

import "slices"

func maxEvents(events [][]int) int {
	//贪心 先按开始天数排序 然后按照持续时间排序
	slices.SortedFunc(events, func(o1 []int, o2 []int) int {
		if o1[0] < o2[0] {
			return -1
		}
		if o1[0] == o2[0] {
			return (o1[0] - o1[1]) - (o2[0] - o2[1])
		}
		return 1
	},
	)
	res := 0
	days := -1
	for i := 0; i < len(events[0]); i++ {
		if events[i][0] >= days && events[i][0] <= days {
			res++
			days = min()
		}
	}

}
