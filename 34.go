package LeetCode

func searchRange(nums []int, target int) []int {
	//思路是找到一个target值 然后往俩边括张

	length := len(nums)
	if length == 1 && nums[0] == target {
		return []int{0, 0}
	}
	left := 0
	right := length - 1
	index := -1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			index = mid
			break
		}
	}
	if index == -1 {
		return []int{-1, -1}
	}
	l := index
	r := index
	for l >= 0 && nums[l] == target {
		l--
	}
	for r < length && nums[r] == target {
		r++
	}
	return []int{l + 1, r - 1}
}
