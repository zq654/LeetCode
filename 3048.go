package LeetCode

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
	//统计从哪一刻开始 changeIndices里面有全部数字 并且index值要大于sum（nums）+len（nums）
	sums := 0
	for _, v := range nums {
		sums += v
	}

	recordMap := make(map[int]bool)
	difCount := 0 //可以封顶的数量
	l := -1
	for index, v := range changeIndices {
		if !recordMap[v] && index >= nums[v-1]-1 {
			recordMap[v] = true
			difCount++
		}
		if difCount == len(nums) && index >= sums+len(nums)-1 {
			l = index
			break
		}
	}
	if l == -1 {
		return -1
	}
	//需要考虑封顶要花1秒时间
	check := func(time int) bool {
		cappingMap := make(map[int]bool) //记录某个值有没有被封顶
		var dfs func(count int, curTime int, mustSpendTime int) bool
		//dfs count表示封顶的数量
		dfs = func(count int, curTime int, mustSpendTime int) bool {
			if count >= len(nums) {
				return true
			}
			if curTime > time || mustSpendTime > time {
				return false
			}
			if curTime-mustSpendTime >= nums[changeIndices[curTime]-1] && !cappingMap[changeIndices[curTime]-1] {
				cappingMap[changeIndices[curTime]-1] = true
				flag1 := dfs(count+1, curTime+1, mustSpendTime+nums[changeIndices[curTime]-1]+1)
				cappingMap[changeIndices[curTime]-1] = false
				if flag1 {
					return true
				}
				flag2 := dfs(count, curTime+1, mustSpendTime)
				return flag1 || flag2
			} else {
				return dfs(count, curTime+1, mustSpendTime)
			}
		}
		return dfs(0, 0, 0)
	}
	r := len(changeIndices) - 1
	for l <= r {
		mid := l + (r-l)>>2
		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l + 1
}
