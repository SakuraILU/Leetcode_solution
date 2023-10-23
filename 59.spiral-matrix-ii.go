/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

// @lc code=start

type Dir struct {
	i, j int
}

var dirs = []Dir{Dir{0, 1}, Dir{1, 0}, Dir{0, -1}, Dir{-1, 0}}

var matrix [][]int
var visited [][]bool
var num int

func generateMatrix(n int) [][]int {
	matrix = make([][]int, 0)
	visited = make([][]bool, 0)
	num = 1
	for i := 0; i < n; i++ {
		row1 := make([]int, n)
		matrix = append(matrix, row1)

		row2 := make([]bool, n)
		visited = append(visited, row2)
	}

	dfs(0, 0, 0, 1)

	return matrix
}

func dfs(i, j, diridx int, num int) bool {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return false
	}

	if visited[i][j] {
		return false
	}
	visited[i][j] = true

	matrix[i][j] = num

	for k := 0; k < 4; k++ {
		ni := i + dirs[diridx].i
		nj := j + dirs[diridx].j
		if dfs(ni, nj, diridx, num+1) {
			return true
		}

		diridx = (diridx + 1) % len(dirs)
	}

	return false
}

// @lc code=end
