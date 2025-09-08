package LeetCode

func findSpecialInteger(arr []int) int {
	//如果超过25% 这这个数字在整个数组里面的长度肯定大于1/4 即分成四分 五个节点里面肯定有一个节点出现两次
	length := len(arr)

	distinctMap := make(map[int]bool)
	for i := 0; i <= 4; i++ {
		value := arr[(i*length/4)%(length-1)]
		if distinctMap[value] {
			return value
		}
		distinctMap[value] = true
	}
	return -1
}
