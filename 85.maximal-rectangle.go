/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */

// @lc code=start
package leetcodesolution

var _matrix [][]byte

func maximalRectangle(matrix [][]byte) int {
	_matrix = matrix
}

func dp(i, j int) (mx, my int) {
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
