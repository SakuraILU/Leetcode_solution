/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
// package leetcodesolution
type State struct {
	i, j int
}

var memo map[State]int

var _matrix [][]byte

func maximalSquare(matrix [][]byte) int {
	_matrix = matrix
	maxlen := 0

	memo = make(map[State]int, 0)
	for i := 0; i < len(_matrix); i++ {
		for j := 0; j < len(_matrix[0]); j++ {
			maxlen = max(maxlen, dp(i, j))
		}
	}

	return maxlen * maxlen
}

func dp(i, j int) (mlen int) {
	if i == 0 || j == 0 {
		if _matrix[i][j] == '1' {
			return 1
		} else {
			return 0
		}
	}

	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	if _matrix[i][j] == '1' {
		mlen = 1 + min(min(dp(i-1, j-1), dp(i-1, j)), dp(i, j-1))
	} else {
		mlen = 0
	}

	memo[state] = mlen
	return mlen
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
