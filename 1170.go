package LeetCode

import "slices"

func numSmallerByFrequency(queries []string, words []string) []int {
	//第一步统计最小的字符的频率
	freArr1 := make([]int, len(queries))
	freArr2 := make([]int, len(words))
	for i := 0; i < len(freArr1); i++ {
		freArr1[i] = getLeastFrequency(queries[i])
	}

	for i := 0; i < len(freArr2); i++ {
		freArr2[i] = getLeastFrequency(words[i])
	}

	slices.Sort(freArr2)
	for i := 0; i < len(freArr1); i++ {
		//在freArr2中查找大于target的最小值
		freArr1[i] = len(freArr2) - findTarget(freArr2, freArr1[i])
	}
	return freArr1
}
func getLeastFrequency(str string) int {
	count := 0
	minByte := 'z'
	byteArr := []rune(str)
	for i := 0; i < len(str); i++ {
		if byteArr[i] < minByte {
			minByte = byteArr[i]
			count = 0
		}
		if byteArr[i] == minByte {
			count++
		}
	}
	return count
}
func findTarget(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
