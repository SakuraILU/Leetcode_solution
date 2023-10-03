/*
 * @lc app=leetcode id=896 lang=golang
 *
 * [896] Monotonic Array
 */

// @lc code=start
var _nums []int

func isMonotonic(nums []int) bool {
	_nums = nums

	return dfs(0, true) || dfs(0, false)
}

func dfs(i int, increase bool) bool {
	if i == len(_nums)-1 {
		return true
	}

	if increase && _nums[i] >= _nums[i+1] {
		return dfs(i+1, true)
	} else if !increase && _nums[i] <= _nums[i+1] {
		return dfs(i+1, false)
	}

	return false
}

// @lc code=end
