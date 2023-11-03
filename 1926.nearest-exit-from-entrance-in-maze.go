/*
 * @lc app=leetcode id=1926 lang=golang
 *
 * [1926] Nearest Exit from Entrance in Maze
 */

// @lc code=start
type Coord struct {
	i, j int
}

type Dir struct {
	i, j int
}

var dirs = []Dir{Dir{1, 0}, Dir{0, 1}, Dir{-1, 0}, Dir{0, -1}}

var _maze [][]byte

const (
	ROAD        = '.'
	WALL        = '+'
	VISITEDROAD = 'x'
)

func nearestExit(maze [][]byte, entrance []int) int {
	_maze = maze

	queue := make([]Coord, 0)
	nqueue := make([]Coord, 0)

	queue = append(queue, Coord{i: entrance[0], j: entrance[1]})
	maze[entrance[0]][entrance[1]] = VISITEDROAD

	length := 0
	for len(queue) > 0 {
		length++
		for _, coord := range queue {
			for _, dir := range dirs {
				ni, nj := coord.i+dir.i, coord.j+dir.j

				if ni < 0 || ni >= len(maze) || nj < 0 || nj >= len(maze[0]) {
					continue
				}
				if maze[ni][nj] == WALL || maze[ni][nj] == VISITEDROAD {
					continue
				}

				maze[ni][nj] = VISITEDROAD

				if isExit(ni, nj) {
					return length
				}

				nqueue = append(nqueue, Coord{i: ni, j: nj})
			}
		}

		queue = nqueue
		nqueue = make([]Coord, 0)
	}

	return -1
}

func isExit(i, j int) bool {
	return i == 0 || i == len(_maze)-1 || j == 0 || j == len(_maze[0])-1
}

// @lc code=end
