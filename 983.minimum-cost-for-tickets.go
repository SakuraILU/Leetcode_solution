/*
 * @lc app=leetcode id=983 lang=golang
 *
 * [983] Minimum Cost For Tickets
 */

// @lc code=start
type State struct {
	i, remaindays int
}

var _days, _costs []int
var _durations = []int{1, 7, 30}
var memo map[State]int

func mincostTickets(days []int, costs []int) int {
	_days, _costs = days, costs
	memo = make(map[State]int)
	return dp(0, 0)
}

func dp(i int, remaindays int) (mcost int) {
	state := State{i: i, remaindays: remaindays}
	if v, ok := memo[state]; ok {
		return v
	}

	mcost = math.MaxInt
	if remaindays > 0 {
		if i < len(_days)-1 {
			mcost = dp(i+1, remaindays-(_days[i+1]-_days[i]))
		} else {
			mcost = 0
		}
	} else {
		for k, cost := range _costs {
			mcost = min(mcost, cost+dp(i, _durations[k]))
		}
	}

	memo[state] = mcost
	return
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
