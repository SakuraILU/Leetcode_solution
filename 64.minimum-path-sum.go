/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
type State struct {
	i, j int
}

var memo map[State]int

var _grid [][]int

func minPathSum(grid [][]int) int {
	_grid = grid
	memo = make(map[State]int, 0)
	return dp(0, 0)
}

func dp(i, j int) (msum int) {
	if i == len(_grid)-1 && j == len(_grid[0])-1 {
		return _grid[i][j]
	} else if i >= len(_grid) || j >= len(_grid[0]) {
		return math.MaxInt
	}

	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}
	msum = _grid[i][j] + min(dp(i+1, j), dp(i, j+1))

	memo[state] = msum
	return msum
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
