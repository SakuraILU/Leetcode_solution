/*
 * @lc app=leetcode id=1306 lang=golang
 *
 * [1306] Jump Game III
 */

// @lc code=start

var garr []int

var memo map[int]bool

func canReach(arr []int, start int) bool {
	garr = arr
	memo = make(map[int]bool, 0)
	return dp(start)
}

func dp(pos int) bool {
	if garr[pos] == 0 {
		return true
	}

	if v, ok := memo[pos]; ok {
		return v
	}
	memo[pos] = false

	jumpstep := garr[pos]
	if pos-jumpstep >= 0 && dp(pos-jumpstep) {
		memo[pos] = true
	} else if pos+jumpstep < len(garr) && dp(pos+jumpstep) {
		memo[pos] = true
	}

	return memo[pos]
}

// @lc code=end
