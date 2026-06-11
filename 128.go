package LeetCode

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}
	res := 0
	for v, _ := range set {
		if _, ok := set[v-1]; ok {
			continue
		}
		count := 1

		for _, exist := set[v+1]; exist; _, exist = set[v+1] {
			count++
			v++
		}
		res = max(res, count)
	}
	return res
}

// 自己想的更新map边界的方法
func longestConsecutive2(nums []int) int {
	recordMap := make(map[int]int)
	res := 0
	for _, num := range nums {
		if recordMap[num] > 0 {
			continue
		}
		count := recordMap[num-1] + recordMap[num+1] + 1
		recordMap[num] = count
		//更新边界
		recordMap[num-recordMap[num-1]] = count
		recordMap[num+recordMap[num+1]] = count
		res = max(res, count)
	}
	return res
}
