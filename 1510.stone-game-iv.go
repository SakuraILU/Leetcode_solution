/*
 * @lc app=leetcode id=1510 lang=golang
 *
 * [1510] Stone Game IV
 */

// @lc code=start
import "fmt"

var memo map[int]bool

func winnerSquareGame(n int) bool {
	memo = make(map[int]bool, 0)
	return dp(n)
}

func dp(n int) bool {
	if v, ok := memo[n]; ok {
		return v
	}

	for k := 1; k*k <= n; k++ {
		if n == k*k || !dp(n-k*k) {
			memo[n] = true
			return true
		}
	}
	memo[n] = false
	return false
}

// @lc code=end
