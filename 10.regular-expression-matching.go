/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
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

func dp(i, j int) (matched bool) {
	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	if i >= len(_s) && j >= len(_p) {
		memo[state] = true
		return memo[state]
	} else if j >= len(_p) {
		memo[state] = false
		return memo[state]
	} else if i == len(_s) {
		// expect p[j: len(p)] can equal to ""...
		// like a*b*c*d*, all the *'s elimate previous letter
		if j < len(_p)-1 && _p[j+1] == '*' {
			memo[state] = dp(i, j+2)
			return memo[state]
		} else {
			memo[state] = false
			return memo[state]
		}
	}

	matched = false
	if j < len(_p)-1 && _p[j+1] == '*' {
		if _s[i] == _p[j] || _p[j] == '.' {
			matched = dp(i+1, j) || dp(i+1, j+2) || dp(i, j+2)
		} else {
			matched = dp(i, j+2)
		}
	} else if _s[i] == _p[j] || _p[j] == '.' {
		matched = dp(i+1, j+1)
	}

	memo[state] = matched
	return memo[state]
}

// @lc code=end
