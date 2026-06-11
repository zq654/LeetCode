package LeetCode

func subarraysDivByK(nums []int, k int) int {
	preArr := make([]int, k)
	preArr[0] = 1

	sum := 0
	res := 0
	for _, num := range nums {
		sum = (sum + num%k + k) % k
		res += preArr[sum]
		preArr[sum]++
	}
	return res
}
