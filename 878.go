package LeetCode

import (
	"math"
	"sort"
)

func nthMagicalNumber(n int, a int, b int) int {
	mods := 1000000000 + 7
	check := func(mx int) bool {
		count := 0
		count += mx / a
		count += mx / b
		//找最小公倍数
		count -= mx / LCM(a, b)
		return count >= n
	}
	res := sort.Search(math.MaxInt64, check)
	return res % mods
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
