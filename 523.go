package LeetCode

func checkSubarraySum(nums []int, k int) bool {
	//要是k的倍速 则x%k 值相同就可以了
	recordMap := make(map[int]int) //value -> index
	recordMap[0] = 0
	sum := 0
	for i, num := range nums {
		sum = (sum + num) % k
		if index, ok := recordMap[sum]; ok && i+1-index >= 2 {
			return true
		}
		if _, ok := recordMap[sum]; !ok {
			recordMap[sum] = i + 1
		}
	}
	return false
}
