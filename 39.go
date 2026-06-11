package LeetCode

import "slices"

// 递归
func combinationSum(candidates []int, target int) (ans [][]int) {
	slices.Sort(candidates)
	path := []int{}
	var dfs func(int, int)
	dfs = func(i, left int) {
		if left == 0 {
			// 找到一个合法组合
			ans = append(ans, slices.Clone(path))
			return
		}

		if i == len(candidates) || left < candidates[i] {
			return
		}

		// 不选
		dfs(i+1, left)

		// 选
		path = append(path, candidates[i])
		dfs(i, left-candidates[i])
		path = path[:len(path)-1] // 恢复现场
	}
	dfs(0, target)
	return ans
}
