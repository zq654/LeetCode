package LeetCode

func solveQueries(nums []int, queries []int) []int {
	//用map 记录value -> index
	recordMap := map[int][]int{}
	for i := 0; i < len(nums); i++ {
		recordMap[nums[i]] = append(recordMap[nums[i]], i)
	}
	//这样map 里面就记录了某个值他的所有index
	for i := 0; i < len(queries); i++ {
		indexArr, ok := recordMap[nums[queries[i]]]
		if !ok || len(indexArr) == 1 {
			queries[i] = -1
		} else {
			index := findValue(indexArr, queries[i]) //index 是在indexArr里面的下标
			var leftDif int
			if index == 0 {
				leftDif = len(nums) - indexArr[len(indexArr)-1] + indexArr[index]
			} else {
				leftDif = queries[i] - indexArr[index-1]
			}
			var rightDif int
			if index == len(indexArr)-1 {
				rightDif = len(nums) - indexArr[index] + indexArr[0]
			} else {
				rightDif = indexArr[index+1] - queries[i]
			}
			queries[i] = min(leftDif, rightDif)
		}
	}
	return queries
}
func findValue(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		}
		if nums[mid] < target {
			l = mid + 1
		}
	}
	return -1
}
