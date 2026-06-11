package LeetCode

func countNicePairs(nums []int) int {
	//nums[i] + rev(nums[j]) == nums[j] + rev(nums[i]) -> nums[i]- rev(nums[i])= nums[j] -rev(nums[j])
	recordMap := make(map[int]int) //value -> count
	res := 0
	for _, num := range nums {
		target := num - getRevValue(num)
		res += recordMap[target]
		recordMap[target]++
	}
	return res % (1e9 + 7)

}
func getRevValue(x int) (res int) {
	for x > 0 {
		res *= 10
		res += x % 10
		x = x / 10
	}
	return res
}
