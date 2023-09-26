/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
type State struct {
	i, j int
}

var memo map[State]int

var gm int
var gn int

func uniquePaths(m int, n int) int {
	gm = m
	gn = n

	memo = make(map[State]int, 0)
	return dp(0, 0)
}

func dp(i, j int) int {
	if i == gm-1 && j == gn-1 {
		return 1
	} else if i >= gm || j >= gn {
		return 0
	}

	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	memo[state] = dp(i+1, j) + dp(i, j+1)
	return memo[state]
}

// @lc code=end
