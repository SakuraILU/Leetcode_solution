/*
 * @lc app=leetcode id=887 lang=golang
 *
 * [887] Super Egg Drop
 */

// @lc code=start
type State struct {
	eggs, n int
}

var memo map[State]int

func superEggDrop(k int, n int) int {
	memo = make(map[State]int, 0)
	return dp(k, n)
}

func dp(eggs int, n int) (mcount int) {
	if n <= 0 {
		return 0
	}

	if eggs == 1 {
		return n
	}

	state := State{eggs: eggs, n: n}
	if v, ok := memo[state]; ok {
		return v
	}

	mcount = math.MaxInt
	for i := 1; i <= n; i++ {
		count := 1 + max(dp(eggs-1, i-1), dp(eggs, n-i))
		mcount = min(mcount, count)
	}

	memo[state] = mcount
	return
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
