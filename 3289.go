package LeetCode

type bitMap struct {
	data []byte
	size int
}

func getSneakyNumbers(nums []int) []int {
	//用位图写一个
	b := initBitMap(100)
	res := make([]int, 0)
	for _, num := range nums {
		if b.checkValueExist(num) {
			res = append(res, num)
		} else {
			b.setValue(num)
		}
	}
	return res
}
func initBitMap(size int) (b *bitMap) {
	b = &bitMap{}
	b.size = (size + 7) / 8
	b.data = make([]byte, b.size)
	return
}

func (b *bitMap) setValue(value int) {
	byteIndex := value / 8
	bitIndex := value % 8
	b.data[byteIndex] |= 1 << bitIndex
}

func (b *bitMap) checkValueExist(value int) bool {
	byteIndex := value / 8
	bitIndex := value % 8
	return b.data[byteIndex]&(1<<bitIndex) != 0
}
