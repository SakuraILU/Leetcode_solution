/*
 * @lc app=leetcode id=799 lang=golang
 *
 * [799] Champagne Tower
 */

// @lc code=start

type State struct {
	i, j int
}

var memo map[State]float64

var _poured int

func champagneTower(poured int, query_row int, query_glass int) float64 {
	_poured = poured

	memo = make(map[State]float64, 0)
	return min(dp(query_row, query_glass), 1)
}

func dp(i, j int) (size float64) {
	if j < 0 || j > i {
		return 0
	}
	if i == 0 {
		return float64(_poured)
	}

	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	mlen := max((dp(i-1, j-1)-1)/2, 0) + max((dp(i-1, j)-1)/2, 0)
	memo[state] = mlen
	return mlen
}

func max(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
