/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */

// @lc code=start
type State struct {
	i, j int
}

var memo map[State]bool

var _s, _p string

func isMatch(s string, p string) bool {
	_s, _p = s, p
	memo = make(map[State]bool)
	return dp(0, 0)
}

func dp(i, j int) bool {
	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	if i >= len(_s) && j >= len(_p) {
		memo[state] = true
		return true
	} else if i >= len(_s) {
		for k := j; k < len(_p); k++ {
			if _p[k] != '*' {
				memo[state] = false
				return false
			}
		}
		memo[state] = true
		return true
	} else if j >= len(_p) {
		memo[state] = false
		return false
	}

	if _p[j] == '*' {
		memo[state] = dp(i+1, j) || dp(i+1, j+1) || dp(i, j+1)
	} else if _p[j] == '?' || _s[i] == _p[j] {
		memo[state] = dp(i+1, j+1)
	}

	return memo[state]
}

// @lc code=end
