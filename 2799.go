package LeetCode

func countCompleteSubarrays(nums []int) int {
	elemMap := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		elemMap[nums[i]] = struct{}{}
	}
	eleCount := len(elemMap)
	numMap := make(map[int]int)
	dif := 0
	l := 0
	res := 0
	for r := 0; r < len(nums); r++ {
		numMap[nums[r]]++
		if numMap[nums[r]] == 1 {
			dif++
		}
		for dif >= eleCount {
			numMap[nums[l]]--
			if numMap[nums[l]] == 0 {
				dif--
			}
			l++
		}
		res += l
	}
	return res
}
