/*
 * @lc app=leetcode id=877 lang=golang
 *
 * [877] Stone Game
 */

// @lc code=start

type State struct {
	i, j int
}

var memo map[State]int

var gpiles []int

func stoneGame(piles []int) bool {
	gpiles = piles

	memo = make(map[State]int, 0)
	return dp(0, len(piles)-1) > 0
}

func dp(i, j int) (mscore int) {
	if i == j {
		return gpiles[i]
	}

	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	mscore = max(gpiles[i]-dp(i+1, j),
		gpiles[j]-dp(i, j-1))

	memo[state] = mscore
	return
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
