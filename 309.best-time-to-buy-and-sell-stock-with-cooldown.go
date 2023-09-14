/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */

// @lc code=start
type State struct {
	day  int
	hold bool
}

var memo map[State]int

var gprices []int

func maxProfit(prices []int) int {
	gprices = prices

	memo = make(map[State]int, 0)
	return dp(0, false)
}

func dp(day int, hold bool) (mprofit int) {
	if day > len(gprices)-1 {
		return 0
	}

	state := State{day: day, hold: hold}
	if v, ok := memo[state]; ok {
		return v
	}

	mprofit = 0
	if !hold {
		mprofit = max(-gprices[day]+dp(day+1, true),
			dp(day+1, false))
	} else {
		mprofit = max(gprices[day]+dp(day+2, false),
			dp(day+1, true))
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

