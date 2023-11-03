/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
var memo map[int]int

func integerBreak(n int) int {
	memo = make(map[int]int)

	// we must split at the first time...
	mmult := 0
	for i := 1; i < n; i++ {
		mmult = max(mmult, i*dp(n-i))
	}

	return mmult
}

func dp(n int) (mmult int) {
	if v, ok := memo[n]; ok {
		return v
	}

	mmult = n // no split
	for i := 1; i < n; i++ {
		mmult = max(mmult, i*dp(n-i))
	}

	memo[n] = mmult
	return mmult
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
