/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start
var _board [][]byte

func solve(board [][]byte) {
	_board = board

	rows, cols := len(_board), len(_board[0])
	for i := 0; i < rows; i++ {
		if _board[i][0] == 'O' {
			dfs(i, 0)
		}
		if _board[i][cols-1] == 'O' {
			dfs(i, cols-1)
		}
	}
	for j := 0; j < cols; j++ {
		if _board[0][j] == 'O' {
			dfs(0, j)
		}
		if _board[rows-1][j] == 'O' {
			dfs(rows-1, j)
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if _board[i][j] == 'Y' {
				_board[i][j] = 'O'
			} else if _board[i][j] == 'O' {
				_board[i][j] = 'X'
			}
		}
	}
}

func dfs(i, j int) {
	if i < 0 || i >= len(_board) || j < 0 || j >= len(_board[0]) {
		return
	}
	if _board[i][j] == 'X' || _board[i][j] == 'Y' {
		return
	}
	_board[i][j] = 'Y'

	dfs(i, j-1)
	dfs(i, j+1)
	dfs(i-1, j)
	dfs(i+1, j)
}

// @lc code=end
