package LeetCode

func numSubarraysWithSum(nums []int, goal int) int {
	count1 := 0
	count2 := 0
	l1 := 0
	l2 := 0
	res1 := 0
	res2 := 0
	//求大于k的数组的数量
	//求大于k+1的数组的数量
	//俩者相减就是==k的数组的数量
	for r := 0; r < len(nums); r++ {
		count1 += nums[r]
		count2 += nums[r]
		for count1 >= goal && l1 <= r {
			count1 -= nums[l1]
			l1++
		}
		for count2 > goal {
			count2 -= nums[l2]
			l2++
		}
		res1 += l1
		res2 += l2
	}
	return res1 - res2
}
