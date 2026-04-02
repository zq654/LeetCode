package LeetCode

type NumArray []int

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	return nums
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return (*this)[right]
	}
	return (*this)[right] - (*this)[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
