/*
 * @lc app=leetcode id=1267 lang=golang
 *
 * [1267] Count Servers that Communicate
 */

// @lc code=start
const (
	EMPTY  = 0
	SERVER = 1
)

func countServers(grid [][]int) (cnt int) {
	rowcnt := make(map[int]int)
	colcnt := make(map[int]int)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == SERVER {
				rowcnt[i]++
				colcnt[j]++
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == SERVER {
				if rowcnt[i] > 1 || colcnt[j] > 1 {
					cnt++
				}
			}
		}
	}

	return cnt
}

// @lc code=end
