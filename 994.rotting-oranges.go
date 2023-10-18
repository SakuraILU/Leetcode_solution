/*
 * @lc app=leetcode id=994 lang=golang
 *
 * [994] Rotting Oranges
 */

// @lc code=start
// package leetcodesolution

type Node struct {
	i, j int
}

type Direction struct {
	x, y int
}

var dirs []Direction = []Direction{Direction{1, 0}, Direction{-1, 0}, Direction{0, 1}, Direction{0, -1}}

func orangesRotting(grid [][]int) (time int) {
	queue := make([]Node, 0)
	nqueue := make([]Node, 0)

	nfresh := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, Node{i: i, j: j})
			} else if grid[i][j] == 1 {
				nfresh++
			}
		}
	}

	for len(queue) > 0 {
		for _, node := range queue {
			for _, dir := range dirs {
				ni := node.i + dir.x
				nj := node.j + dir.y

				if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[0]) {
					continue
				}

				if grid[ni][nj] == 1 {
					grid[ni][nj] = 2
					nfresh--
					nqueue = append(nqueue, Node{i: ni, j: nj})
				}
			}
		}

		if len(nqueue) > 0 {
			time++
		}

		queue = nqueue
		nqueue = make([]Node, 0)
	}

	if nfresh > 0 {
		time = -1
	}

	return
}

// @lc code=end
