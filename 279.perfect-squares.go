/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start
var memo map[int]int

func numSquares(n int) int {
	memo = make(map[int]int)
	return dp(n)
}

func dp(n int) (mnum int) {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	if v, ok := memo[n]; ok {
		return v
	}
	mnum = math.MaxInt
	for i := n / 2; i > 0; i-- {
		if n-i*i >= 0 {
			mnum = min(mnum, 1+dp(n-i*i))
		}
	}

	memo[n] = mnum
	return mnum
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
