/*
 * @lc app=leetcode id=1884 lang=golang
 *
 * [1884] Egg Drop With 2 Eggs and N Floors
 */

// @lc code=start

import (
	"math"
)

type State struct {
	l, h, eggs int
}

var memo map[State]int

func twoEggDrop(n int) int {
	memo = make(map[State]int, 0)

	return dp(State{l: 1, h: n, eggs: 2})
}

func dp(state State) (mmove int) {
	if v, ok := memo[state]; ok {
		return v
	}

	if state.l > state.h {
		return 0
	} else if state.l == state.h {
		return 1
	}

	if state.eggs == 1 {
		return state.h - state.l + 1
	}

	mmove = math.MaxInt
	for i := state.l; i <= state.h; i++ {
		move := 1 + max(dp(State{l: state.l, h: i - 1, eggs: state.eggs - 1}),
			dp(State{l: i + 1, h: state.h, eggs: state.eggs}))
		mmove = min(mmove, move)
	}

	memo[state] = mmove
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
