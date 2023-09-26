/*
 * @lc app=leetcode id=980 lang=golang
 *
 * [980] Unique Paths III
 */

// @lc code=start

type State struct {
	i, j int
}

var choices = []State{State{0, 1}, State{0, -1}, State{1, 0}, State{-1, 0}}

var ggrid [][]int
var visited [][]bool

var targetstep int = 0 // how many steps robot need to move (all 0's and 2)
var curstep int = 0

func uniquePathsIII(grid [][]int) int {
	ggrid = grid

	visited = make([][]bool, len(ggrid))
	for i := 0; i < len(ggrid); i++ {
		visited[i] = make([]bool, len(ggrid[0]))
	}

	x, y := -1, -1
	curstep, targetstep = 0, 0
	for i := 0; i < len(ggrid); i++ {
		for j := 0; j < len(ggrid[0]); j++ {
			if ggrid[i][j] == 1 {
				x, y = i, j
				visited[x][y] = true
			} else if ggrid[i][j] == 0 || ggrid[i][j] == 2 {
				targetstep++
			}
		}
	}
	return dp(x, y)
}

func dp(i, j int) (num int) {
	if ggrid[i][j] == -1 {
		return 0
	} else if ggrid[i][j] == 2 {
		if curstep == targetstep {
			return 1
		} else {
			return 0
		}
	}

	for _, choice := range choices {
		ni := i + choice.i
		nj := j + choice.j

		// not out of bound
		if ni < 0 || ni >= len(ggrid) || nj < 0 || nj >= len(ggrid[0]) {
			continue
		}
		// not visited
		if visited[ni][nj] {
			continue
		}

		visited[ni][nj] = true
		curstep++
		num += dp(ni, nj)
		curstep--
		visited[ni][nj] = false

	}

	return num
}

// @lc code=end
