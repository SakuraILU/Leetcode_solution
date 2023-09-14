/*
 * @lc app=leetcode id=375 lang=golang
 *
 * [375] Guess Number Higher or Lower II
 */

// @lc code=start

import "math"

type Pair struct {
	i, j interface{}
}

var memo = make(map[Pair]int, 0)

func getMoneyAmount(n int) int {
	return dp(1, n)
}

func dp(i, j int) int {
	if i >= j {
		return 0
	}

	if v, ok := memo[Pair{i, j}]; ok {
		return v
	}

	minamount := math.MaxInt
	for k := i; k <= j; k++ {
		minamount = min(minamount, k+max(dp(i, k-1), dp(k+1, j)))
	}

	memo[Pair{i, j}] = minamount
	return minamount
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

// @lc code=end
