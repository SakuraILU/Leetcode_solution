/*
 * @lc app=leetcode id=123 lang=golang
 *
 * [123] Best Time to Buy and Sell Stock III
 */

// @lc code=start
type State struct {
	day       int
	hold      bool
	transtime int
}

var memo map[State]int

var gprices []int

func maxProfit(prices []int) int {
	gprices = prices
	memo = make(map[State]int, 0)
	return dp(0, false, 0)
}

func dp(day int, hold bool, transtime int) (mprofit int) {
	if transtime >= 2 {
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
