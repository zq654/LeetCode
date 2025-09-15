package LeetCode

func splitArray(nums []int, k int) int {
	//前缀和 + 差分 + 二分
	length := len(nums)
	l := nums[0]
	r := nums[0]
	for i := 1; i < length; i++ {
		l = max(l, nums[i])
		r += nums[i]
		nums[i] = r
	}
	if k <= 1 {
		return r
	}
	//贪心 每次求和尽可能小于等于maxValue 若最后分割出来的长度大于k则return false
	check := func(maxValue int) bool {
		preSum := 0
		count := 0
		for i := 0; i < length; {
			if nums[i]-preSum > maxValue {
				count++
				preSum = nums[i-1]
				continue
			}
			if i == length-1 && nums[i]-preSum > 0 {
				count++
			}
			if count > k {
				return false
			}
			i++
		}
		return true
	}
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l - 1
}
