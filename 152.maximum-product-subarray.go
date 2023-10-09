/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
type State struct {
	i     int
	ismax bool
}

var memo map[State]int

var _nums []int

func maxProduct(nums []int) int {
	_nums = nums
	memo = make(map[State]int)

	mmul := math.MinInt
	for i := 0; i < len(nums); i++ {
		mmul = max(mmul, dp(i, true))
	}

	return mmul
}

func dp(i int, ismax bool) int {
	if i == 0 {
		return _nums[i]
	}

	state := State{i: i, ismax: ismax}
	if v, ok := memo[state]; ok {
		return v
	}

	if ismax {
		if _nums[i] >= 0 {
			memo[state] = max(_nums[i], _nums[i]*dp(i-1, true))
		} else if _nums[i] < 0 {
			memo[state] = max(_nums[i], _nums[i]*dp(i-1, false))
		}
	} else {
		if _nums[i] >= 0 {
			memo[state] = min(_nums[i], _nums[i]*dp(i-1, false))
		} else if _nums[i] < 0 {
			memo[state] = min(_nums[i], _nums[i]*dp(i-1, true))
		}
	}

	return memo[state]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
