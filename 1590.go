package LeetCode

// 用hash表
func minSubarray(nums []int, p int) int {
	//移除某个数组之后使其可以被p整除
	//转化下题意就是 移除的数组和的余数等同与这个数组的余数
	//提示里面指定nums[i]是大于1的 因此前缀和单调递增
	pre := make(map[int]int) //不能用数组 用数组会oom
	sum := 0
	//先获取整个数组的余数
	count := 0
	for _, num := range nums {
		count += num
	}
	k := count % p
	if k == 0 {
		return 0
	}
	res := len(nums)
	pre[0] = -1 //前缀和为0的初始化

	//根据前缀和求答案
	for i, num := range nums {
		sum = (sum + num) % p
		index, ok := pre[(sum-k+p)%p]
		if ok {
			res = min(res, i-index)
		}
		pre[sum] = i
	}
	if res == len(nums) {
		return -1
	}
	return res
}
