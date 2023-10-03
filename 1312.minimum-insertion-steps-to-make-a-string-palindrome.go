/*
 * @lc app=leetcode id=1312 lang=golang
 *
 * [1312] Minimum Insertion Steps to Make a String Palindrome
 */

// @lc code=start
type State struct {
	i, j int
}

var memo map[State]int

var _s string

func minInsertions(s string) int {
	_s = s

	memo = make(map[State]int, 0)
	return dp(0, len(s)-1)
}

func dp(i, j int) (mcnt int) {
	if i > j {
		return 0
	}

	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	mcnt = math.MaxInt
	if _s[i] == _s[j] {
		mcnt = dp(i+1, j-1)
	}

	mcnt = min(min(mcnt, 1+dp(i+1, j)), 1+dp(i, j-1))

	memo[state] = mcnt
	return mcnt
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
