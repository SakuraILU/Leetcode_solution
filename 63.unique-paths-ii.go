/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start
type State struct {
	i, j int
}

var memo map[State]int

var gobstacleGrid [][]int
var rows int
var cols int

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	gobstacleGrid = obstacleGrid
	rows = len(obstacleGrid)
	cols = len(obstacleGrid[0])

	memo = make(map[State]int, 0)

	return dp(0, 0)
}

func dp(i, j int) (num int) {
	if i >= rows || j >= cols {
		return 0
	} else if gobstacleGrid[i][j] == 1 {
		return 0
	} else if i == rows-1 && j == cols-1 {
		return 1
	}

	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	num += dp(i+1, j) + dp(i, j+1)

	memo[state] = num
	return num
}

// @lc code=end

