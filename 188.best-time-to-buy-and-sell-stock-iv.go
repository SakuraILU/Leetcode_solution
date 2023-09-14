/*
 * @lc app=leetcode id=188 lang=golang
 *
 * [188] Best Time to Buy and Sell Stock IV
 */

// @lc code=start
type State struct {
	day       int
	hold      bool
	transtime int
}

var memo map[State]int

var gprices []int
var gk int

func maxProfit(k int, prices []int) int {
	gprices = prices
	gk = k

	memo = make(map[State]int, 0)
	return dp(0, false, 0)
}

func dp(day int, hold bool, transtime int) (mprofit int) {
	if transtime >= gk {
		return 0
	}
	if day > len(gprices)-1 {
		return 0
	}

	state := State{day: day, hold: hold, transtime: transtime}
	if v, ok := memo[state]; ok {
		return v
	}

	if !hold {
		mprofit = max(-gprices[day]+dp(day+1, true, transtime),
			dp(day+1, false, transtime))
	} else {
		mprofit = max(gprices[day]+dp(day+1, false, transtime+1),
			dp(day+1, true, transtime))
	}

	memo[state] = mprofit
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

