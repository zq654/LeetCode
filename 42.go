package LeetCode

func trap(height []int) int {
	//这个格子能接的水取决与 min(左边最大,右测最大）- height[i]
	res := 0
	l := 0
	r := len(height) - 1
	lMaxValue := 0
	rMaxValue := 0
	for l <= r {
		//左侧最高值小于右侧最高值
		if lMaxValue < rMaxValue {
			lMaxValue = max(height[l], lMaxValue)
			res += lMaxValue - height[l]
			l++
		} else {
			rMaxValue = max(height[r], rMaxValue)
			res += rMaxValue - height[r]
			r--
		}
	}
	return res
}
