package LeetCode

import "slices"

func maxProfit(inventory []int, orders int) int {
	// 二分查找卖到最后最小值
	check := func(v int) bool {
		count := 0
		for _, num := range inventory {
			if num >= v {
				count += num - v + 1
				if count >= orders {
					return true
				}
			}
		}
		return false
	}

	l, r := 0, slices.Max(inventory)
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	res := 0
	for _, num := range inventory {
		// 计算从inventory[i]降到r+1能卖出的数量
		fullCount := num - l + 1
		if fullCount <= 0 {
			continue
		}

		if fullCount <= orders {
			// 订单充足，可将该库存全部卖到l
			profit := (num + l) * fullCount / 2
			res = (res + profit) % 1000000007
			orders -= fullCount
		}
	}

	// 处理剩余订单
	res = (res + orders*(l-1)) % 1000000007
	return res
}
