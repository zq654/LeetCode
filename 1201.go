package LeetCode

import (
	"math"
	"sort"
)

func nthUglyNumber(n int, a int, b int, c int) int {
	check := func(mx int) bool {
		count := 0
		count += mx / a
		count += mx / b
		count += mx / c
		count -= mx / LCM(a, b)
		count -= mx / LCM(a, c)
		count -= mx / LCM(c, b)
		count += mx / LCM(LCM(a, b), c)
		return count >= n
	}
	return sort.Search(math.MaxInt64, check)
}

// GCD 用欧几里得算法求两个数的最大公约数
func GCD(a, b int) int {
	for b != 0 {
		// 辗转相除：a = b, b = a % b，直到b=0
		a, b = b, a%b
	}
	return a
}

// LCM 基于GCD求两个数的最小公倍数
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	// 先除后乘，避免a*b导致的整数溢出（比直接a*b/GCD更安全）
	return a / GCD(a, b) * b
}
