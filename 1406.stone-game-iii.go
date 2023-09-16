/*
 * @lc app=leetcode id=1406 lang=golang
 *
 * [1406] Stone Game III
 */

// @lc code=start
// package leetcodesolution

import "fmt"

var memo map[int]int

var gstoneValue []int

func stoneGameIII(stoneValue []int) string {
	gstoneValue = stoneValue
	memo = make(map[int]int, 0)

	res := dp(0)
	if res > 0 {
		return "Alice"
	} else if res < 0 {
		return "Bob"
	} else {
		return "Tie"
	}
}

func dp(i int) (mscore int) {
	// the value of pile may be negative,
	// so...don't take all the piles if 2 or 3 left...
	// but, if only 1 left, must take this according to the rule
	if i == len(gstoneValue)-1 {
		return gstoneValue[i]
	} else if i > len(gstoneValue)-1 {
		return 0
	}

	if v, ok := memo[i]; ok {
		return v
	}

	mscore = math.MinInt
	curscore := 0
	for k := 0; k < 3; k++ {
		if i+k < len(gstoneValue) {
			curscore += gstoneValue[i+k]
			mscore = max(mscore, curscore-dp(i+k+1))
		}
	}

	// fmt.Printf("Return from %v: %v\n", i, mscore)
	memo[i] = mscore
	return mscore
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
