package LeetCode

func maxSum(nums []int, m int, k int) int64 {
	//用map记录 一个num出现的次数
	appearMap := make(map[int]int)
	dif := 0
	count := 0
	res := 0
	for i := 0; i < k; i++ {
		count = count + nums[i]
		if appearMap[nums[i]] == 0 {
			dif++
		}
		appearMap[nums[i]] = appearMap[nums[i]] + 1
	}
	if dif >= m {
		res = count
	}
	r := k
	l := 0
	for r < len(nums) {
		if appearMap[nums[r]] == 0 {
			dif++
		}
		if appearMap[nums[l]] == 1 {
			dif--
		}
		appearMap[nums[r]] = appearMap[nums[r]] + 1
		appearMap[nums[l]] = appearMap[nums[l]] - 1
		count = count + nums[r] - nums[l]
		if dif >= m {
			res = max(res, count)
		}
		l++
		r++
	}
	return int64(res)
}
