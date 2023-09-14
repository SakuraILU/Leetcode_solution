/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start

var gnums []int

var memo map[int]bool

func canJump(nums []int) bool {
	gnums = nums
	memo = make(map[int]bool, 0)
	return dp(0)
}

func dp(pos int) bool {
	if pos+gnums[pos] >= len(gnums)-1 {
		return true
	}

	if v, ok := memo[pos]; ok {
		return v
	}

	memo[pos] = false

	for jumpstep := 1; jumpstep <= gnums[pos]; jumpstep++ {
		if dp(pos + jumpstep) {
			memo[pos] = true
			return true
		}
	}

	return false
}

// @lc code=end
