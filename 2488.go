package LeetCode

import "math"

func countSubarrays(nums []int, k int) int {
	//前缀和 条件里面每个数互不相同
	//因此 当值大于k的时候值+1 值小于k的时候-1
	//并且子序列里面必须要包含k 因此在k左侧的才可以加入map 包含k才可以统计答案
	//当子序列长度为偶数的时候 diff是1
	//当子序列长度是奇数的时候 diff是0
	res := 0
	kIndex := math.MaxInt64
	pre := make(map[int][]int) //sum -> [下标为奇数的数量，下标为偶数的数量]
	pre[0] = []int{1, 0}       //-1也要加入进去
	sum := 0
	for i, num := range nums {
		if num == k {
			kIndex = i
		}
		if num > k {
			sum++
		} else if num < k {
			sum--
		}
		if value, ok := pre[sum-1]; ok && i%2 == 0 && i >= kIndex {
			res += value[1]
		}
		if value, ok := pre[sum]; ok && i%2 == 0 && i >= kIndex {
			res += value[0]
		}
		if value, ok := pre[sum-1]; ok && i%2 == 1 && i >= kIndex {
			res += value[0]
		}
		if value, ok := pre[sum]; ok && i%2 == 1 && i >= kIndex {
			res += value[1]
		}
		if i < kIndex {
			_, ok := pre[sum]
			if !ok {
				pre[sum] = []int{0, 0}
			}
			if i%2 == 1 {
				pre[sum][0]++
			} else {
				pre[sum][1]++
			}
		}
	}
	return res
}

// 优化2.0
// 如果大于k的比小于k的多一个 那么这个子序列他长度一定是 x+x-1+1 =2*x 为偶数 （+1 是因为要加上等于k值的那种值本身）
// 同理大于k和小于k的一样多 他的子序列长度一定是单数
func countSubarrays(nums []int, k int) int {
	//前缀和 条件里面每个数互不相同
	//因此 当值大于k的时候值+1 值小于k的时候-1
	//并且子序列里面必须要包含k 因此在k左侧的才可以加入map 包含k才可以统计答案
	//当子序列长度为偶数的时候 diff是1
	//当子序列长度是奇数的时候 diff是0
	res := 0
	findK := false
	pre := make(map[int]int, len(nums)) //sum -> count
	pre[0] = 1                          //-1也要加入进去
	sum := 0
	for _, num := range nums {
		if num == k {
			findK = true
		}
		if num > k {
			sum++
		} else if num < k {
			sum--
		}
		if findK {
			res += pre[sum] + pre[sum-1]
		} else {
			pre[sum]++
		}
	}
	return res
}
