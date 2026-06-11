package LeetCode

func interchangeableRectangles(rectangles [][]int) int64 {
	recordMap := make(map[float32]int) //value -> count
	count := 0
	for _, rectangle := range rectangles {
		value := float32(rectangle[0]) / float32(rectangle[1])
		count += recordMap[value]
		recordMap[value]++
	}
	return int64(count)
}
