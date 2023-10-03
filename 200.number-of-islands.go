/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
// package leetcodesolution

var _grid [][]byte

func numIslands(grid [][]byte) int {
	_grid = grid

	cnt := 0
	for i := 0; i < len(_grid); i++ {
		for j := 0; j < len(_grid[0]); j++ {
			if _grid[i][j] == '1' {
				dfs(i, j)
				cnt++
			}
		}
	}

	return cnt
}

func dfs(i, j int) {
	if i < 0 || i >= len(_grid) || j < 0 || j >= len(_grid[0]) {
		return
	}

	if _grid[i][j] == '0' {
		return
	}
	_grid[i][j] = '0'

	dfs(i+1, j)
	dfs(i-1, j)
	dfs(i, j+1)
	dfs(i, j-1)
}

// @lc code=end
