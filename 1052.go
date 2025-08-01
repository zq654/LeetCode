package LeetCode

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	count := 0
	res := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			count += customers[i]
		}
	}
	for i := 0; i < len(customers); i++ {

		if grumpy[i] == 1 {
			count += customers[i]
		}
		if i < minutes-1 {
			continue
		}
		res = max(res, count)
		l := i - minutes + 1
		if grumpy[l] == 1 {
			count -= customers[l]
		}
	}
	return res
}
