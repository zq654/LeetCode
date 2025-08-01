package LeetCode

func numberOfSubarrays(nums []int, k int) int {
	res1 := 0
	res2 := 0
	count1 := 0
	count2 := 0
	l1 := 0
	l2 := 0
	for r := 0; r < len(nums); r++ {
		if nums[r]%2 == 1 {
			count1++
			count2++
		}
		for count1 > k {
			if nums[l1]%2 == 1 {
				count1--
			}
			l1++
		}
		res1 += l1
		for count2 >= k {
			if nums[l2]%2 == 1 {
				count2--
			}
			l2++
		}
		res2 += l2
	}
	return res2 - res1
}
