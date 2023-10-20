/*
 * @lc app=leetcode id=289 lang=golang
 *
 * [289] Game of Life
 */

// @lc code=start
const (
	// the settings in the problem description
	DEAD int = 0
	LIVE int = 1
	// personal settings
	LIVE2DEAD int = 2
	DEAD2LIVE int = 3
)

func gameOfLife(board [][]int) {
	// record changes
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			NLiveNeighbors := NumberOfLiveNeighbors(board, i, j)
			fmt.Println(i, j, NLiveNeighbors)

			switch board[i][j] {
			case LIVE:
				if NLiveNeighbors < 2 || NLiveNeighbors > 3 {
					board[i][j] = LIVE2DEAD
				}
			case DEAD:
				if NLiveNeighbors == 3 {
					board[i][j] = DEAD2LIVE
				}
			}
		}
	}

	// fresh state
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == LIVE2DEAD {
				board[i][j] = DEAD
			} else if board[i][j] == DEAD2LIVE {
				board[i][j] = LIVE
			}
		}
	}
}

func NumberOfLiveNeighbors(board [][]int, x, y int) (num int) {
	for i := max(x-1, 0); i < min(x+2, len(board)); i++ {
		for j := max(y-1, 0); j < min(y+2, len(board[0])); j++ {
			// only count neighbors, not itself
			if i == x && j == y {
				continue
			}

			if board[i][j] == LIVE || board[i][j] == LIVE2DEAD {
				num++
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
