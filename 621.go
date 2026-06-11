package LeetCode

import (
	"container/heap"
)

func leastInterval(tasks []byte, n int) int {
	//思路:将n+1看成一个周期
	//维护一个大根堆 每次从大根堆里面取n个数
	//如果堆里面有n个数 则直接时间加n 然后堆里面的数字减一在放回去
	//如果堆里面数字小于n 贪心新开max(task nums)-1 个周期
	//再把底下剩下的几个任务写进去
	n = n + 1
	term := 0 //记录已经发生了多少周期
	res := 0  //统计时间
	bigHeap := &my621Heap{}
	heap.Init(bigHeap)
	recordMap := make([]int, 26)
	for _, task := range tasks {
		recordMap[task-'A']++
	}
	//往堆里面初始化数据
	for i, count := range recordMap {
		if count > 0 {
			heap.Push(bigHeap, item{
				task:  byte('A' + i),
				count: count,
			})
		}
	}
	//如果堆里面有充足的任务够填满一个周期，则取出来填补周期
	for bigHeap.Len() >= n {
		res += n
		term++
		tempItemList := make([]item, 0)
		for i := 0; i < n; i++ {
			popItem := heap.Pop(bigHeap)
			tempItemList = append(tempItemList, popItem.(item))
		}
		for _, itemData := range tempItemList {
			itemData.count--
			if itemData.count > 0 {
				heap.Push(bigHeap, itemData)
			}
		}
	}
	if bigHeap.Len() > 0 {
		//最后查看任务列表里面哪个任务还剩下待处理次数最大
		maxLeftItem := heap.Pop(bigHeap).(item).count
		leftTime := 1
		for bigHeap.Len() > 0 {
			if heap.Pop(bigHeap).(item).count < maxLeftItem {
				break
			}
			leftTime++
		}
		res += (maxLeftItem-1)*n + leftTime
	}
	return res
}

type item struct {
	task  byte
	count int
}

type my621Heap []item

func (h *my621Heap) Less(i, j int) bool {
	return (*h)[i].count > (*h)[j].count
}

func (h *my621Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *my621Heap) Len() int {
	return len(*h)
}

func (h *my621Heap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *my621Heap) Push(v any) {
	*h = append(*h, v.(item))
}
