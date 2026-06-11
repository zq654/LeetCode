package LeetCode

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	mxI, mnI := 0, 0
	l := 0
	for r := indexDifference; r < len(nums); r++ {
		v := nums[r]
		if nums[l] > nums[mxI] {
			mxI = l
		}
		if nums[l] < nums[mnI] {
			mnI = l
		}
		if abs(v-nums[mxI]) >= valueDifference {
			return []int{mxI, r}
		}
		if abs(v-nums[mnI]) >= valueDifference {
			return []int{mnI, r}
		}

		l++
	}
	return []int{-1, -1}

}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
