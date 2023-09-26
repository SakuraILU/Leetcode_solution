/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */

// @lc code=start
var table [][]byte

func totalNQueens(n int) int {
	table = make([][]byte, n)
	for i := 0; i < n; i++ {
		table[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			table[i][j] = '.'
		}
	}

	return dfs(0)
}

func dfs(row int) (num int) {
	if row >= len(table) {
		return 1
	}

	for col := 0; col < len(table[0]); col++ {
		table[row][col] = 'Q'

		if isValid(row, col) {
			num += dfs(row + 1)
		}
		table[row][col] = '.'
	}

	return num
}

func isValid(row, col int) bool {
	for i := 1; i <= row; i++ {
		if table[row-i][col] == 'Q' {
			return false
		}
		if col-i >= 0 && table[row-i][col-i] == 'Q' {
			return false
		}

		if col+i < len(table[0]) && table[row-i][col+i] == 'Q' {
			return false
		}
	}

	return true
}

// @lc code=end
