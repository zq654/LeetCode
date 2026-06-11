package LeetCode

func numOfSubarrays(arr []int) int {
	//奇数加偶数还是为奇数 同理奇数减偶数还是奇数 偶数减奇数也是奇数
	res := 0
	sum := 0
	count1 := 1 //记录前面有几个和偶数数组  为1是因为0也是偶数
	count2 := 0 //记录前面有几个和为奇数的数组
	for _, i := range arr {
		sum += i
		if sum%2 == 1 {
			res += count1
			count2++
		} else {
			res += count2
			count1++
		}
	}
	return res % (1e9 + 7)
}
