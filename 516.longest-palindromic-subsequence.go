/*
 * @lc app=leetcode id=516 lang=golang
 *
 * [516] Longest Palindromic Subsequence
 */

// @lc code=start
type State struct {
	i, j int
}

var memo map[State]int

var _s string

func longestPalindromeSubseq(s string) int {
	_s = s
	memo = make(map[State]int, 0)
	return dp(0, len(s)-1)
}

func dp(i, j int) (mlen int) {
	if i > j {
		return 0
	} else if i == j {
		return 1
	}

	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	if _s[i] == _s[j] {
		mlen = dp(i+1, j-1) + 2
	}

	mlen = max(max(mlen, dp(i, j-1)), dp(i+1, j))

	memo[state] = mlen
	return mlen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
