/*
 * @lc app=leetcode id=115 lang=golang
 *
 * [115] Distinct Subsequences
 */

// @lc code=start
type State struct {
	i, j int
}

var memo map[State]int

var _s, _t string

func numDistinct(s string, t string) int {
	_s, _t = s, t
	memo = make(map[State]int)

	return dp(0, 0)
}

func dp(i, j int) int {
	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	if j >= len(_t) {
		memo[state] = 1
		return 1
	} else if i >= len(_s) {
		memo[state] = 0
		return 0
	}

	if len(_s)-i < len(_t)-j {
		memo[state] = 0
		return 0
	}

	if _s[i] == _t[j] {
		memo[state] = dp(i+1, j+1) + dp(i+1, j)
	} else {
		memo[state] = dp(i+1, j)
	}

	return memo[state]
}

// @lc code=end
