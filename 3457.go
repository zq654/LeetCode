package LeetCode

import (
	"fmt"
	"slices"
)

func maxWeight(pizzas []int) int64 {
	//奇数天得到最大价值的披萨
	//偶数天得到价值第二大的披萨
	//贪心 先处理奇数天 从奇数天里面拿最大的n个 和末尾的3n个垫一下

	days := len(pizzas) / 4
	slices.SortFunc(pizzas, func(a, b int) int {
		return b - a
	})
	fmt.Println(pizzas)
	oddDay := (days + 1) / 2
	evenDay := days / 2
	res := 0
	for i := 0; i < oddDay; i++ {
		res += pizzas[i]
	}
	for i := 0; i < evenDay; i++ {
		res += pizzas[oddDay+i*2+1]
	}
	return int64(res)
}
