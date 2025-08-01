package LeetCode

func searchInsert(nums []int, target int) int {
	//找到小于等于这个的第一个值
	l := 0
	r := len(nums)
	for l < r { //[l,r)
		mid := l + (r-l)/2
		if nums[mid] >= r {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
