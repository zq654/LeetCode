package LeetCode

import "math"

// 自己写法 记录下标
// 贪心策略是 如果有俩个可选值 x1 x2  x1在前;x2在后
// x1到target 的答案可以分成 (x1到x2-1) + (x2到target)
// 如果x1到x2-1的值小于零(也就是 x1到x2 不如直接选择x2)就干脆不要前面那部分 选择x2就行
func maximumSubarraySum(nums []int, k int) int64 {
	//绝对值为k 只要后续有x-k 或x+k 就可以组成一个目标值
	preSumArr := make([]int, len(nums)+1)
	recordMap := make(map[int]int) //target -> index
	res := math.MinInt64
	for i, num := range nums {
		preSumArr[i+1] = preSumArr[i] + num
		if index, ok := recordMap[num]; ok {
			res = max(res, preSumArr[i+1]-preSumArr[index])

		}
		preIndex, ok := recordMap[num+k]
		//不存在 或者 存在且满足贪心条件的时候才更新
		if !ok || (ok && preSumArr[i+1]-preSumArr[preIndex] < num) {
			recordMap[num+k] = i
		}

		preIndex, ok = recordMap[num-k]
		//不存在 或者 存在且满足贪心条件的时候才更新
		if !ok || (ok && preSumArr[i+1]-preSumArr[preIndex] < num) {
			recordMap[num-k] = i
		}
	}
	if res == math.MinInt64 {
		return 0
	}
	return int64(res)
}

// 灵神写法 直接记录前缀和
// 贪心思路为:无论前面怎么选 到当前点的前缀和是不会变的
// 因此选择上一次的前缀和越小越好 这要减数变小 差值就会最大增大
func maximumSubarraySum(nums []int, k int) int64 {
	minS := map[int]int{}
	sum := 0
	ans := math.MinInt
	for _, x := range nums {
		s, ok := minS[x+k]
		if ok {
			ans = max(ans, sum+x-s)
		}

		s, ok = minS[x-k]
		if ok {
			ans = max(ans, sum+x-s)
		}

		s, ok = minS[x]
		if !ok || sum < s {
			minS[x] = sum
		}

		sum += x
	}
	if ans == math.MinInt {
		return 0
	}
	return int64(ans)
}
