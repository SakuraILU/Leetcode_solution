/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
const VISITED int = -200 // -100 <= matrix[i][j] <= 100

type DIR struct {
	di, dj int
}

var dirs = []DIR{DIR{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
var diridx = 0

var _matrix [][]int
var track []int

func spiralOrder(matrix [][]int) []int {
	_matrix = matrix
	track = make([]int, 0)

	dfs(0, 0, 0)

	return track
}

func dfs(i, j int, diridx int) bool {
	if i < 0 || i >= len(_matrix) || j < 0 || j >= len(_matrix[0]) {
		return false
	}
	if _matrix[i][j] == VISITED {
		return false
	}
	track = append(track, _matrix[i][j])

	_matrix[i][j] = VISITED

	for k := 0; k < len(dirs); k++ {
		dir := dirs[diridx]
		ni, nj := i+dir.di, j+dir.dj

		if dfs(ni, nj, diridx) {
			return true
		} else {
			diridx = (diridx + 1) % len(dirs)
		}
	}

	return true
}

// @lc code=end
