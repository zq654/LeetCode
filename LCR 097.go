package LeetCode

func numDistinct(s string, t string) int {
	length := len(t)
	recordArr := make([]int, length)
	recordMap := make(map[byte][]int, 0)
	for i, _ := range t {
		v := t[i]
		recordMap[v] = append(recordMap[v], i)
	}

	recordNumsArr := make(map[byte]int)

	for i, _ := range s {
		v := s[i]
		recordNumsArr[v]++
		for numIndex := len(recordMap[v]) - 1; i >= 0; i-- {
			index := recordMap[v][numIndex]
			if index == 0 {
				recordArr[index] += 1
			} else {
				recordArr[index] += recordArr[index-1]
			}
		}
	}

	return recordArr[length-1]
}
