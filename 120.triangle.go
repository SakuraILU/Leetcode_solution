/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */

// @lc code=start

type State struct {
	i, j int
}

var memo map[State]int

var _triangle [][]int

func minimumTotal(triangle [][]int) int {
	_triangle = triangle

	memo = make(map[State]int)
	return dp(0, 0)
}

func dp(i, j int) int {
	if i == len(_triangle)-1 {
		return _triangle[i][j]
	}

	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	msum := _triangle[i][j] + min(dp(i+1, j), dp(i+1, j+1))

	memo[state] = msum
	return msum
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
