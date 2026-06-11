package LeetCode

import "fmt"

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix[0])
	l := matrix[0][0]
	r := matrix[0][n-1]
	for i := 1; i < n; i++ {
		if matrix[i][n-1] > r {
			r = matrix[i][n-1]
		}
	}
	check := func(mx int) bool {
		count := 0
		for i := 0; i < n; i++ {
			if matrix[i][0] > mx {
				break
			}
			count += findTarget(matrix[i], mx)
		}
		return count >= k
	}
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

// 找到小于等于n的值的index
func findTarget(slices []int, x int) int {
	l := 0
	r := len(slices) - 1
	for l <= r {
		mid := l + (r-l)>>1
		if slices[mid] > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	fmt.Println("findtarget")
	return l
}
