/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

// @lc code=start

import "fmt"

var gs string

var memo map[int]int

func longestValidParentheses(s string) int {
	gs = s
	memo = make(map[int]int, 0)

	mlen := 0
	for i := 0; i < len(s); i++ {
		mlen = max(mlen, dp(i))
		fmt.Println(mlen)
	}
	return mlen
}

func dp(i int) (mlen int) {
	if i <= 0 {
		return 0
	}

	if v, ok := memo[i]; ok {
		return v
	}

	switch gs[i] {
	case '(':
		mlen = 0
	case ')':
		premlen := dp(i - 1)
		start := i - premlen - 1
		if start >= 0 && gs[start] == '(' {
			mlen = premlen + 2 + dp(start-1)
		}
	}

	memo[i] = mlen
	return
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
