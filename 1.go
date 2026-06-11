package LeetCode

// 方法一 用map维护
func twoSum(nums []int, target int) []int {
	recordMap := make(map[int]int) //value -> index
	for index1, v := range nums {
		if index2, ok := recordMap[target-v]; ok {
			return []int{index1, index2}
		} else {
			recordMap[v] = index1
		}
	}
	return []int{}
}
