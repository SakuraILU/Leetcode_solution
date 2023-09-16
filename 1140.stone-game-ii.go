/*
 * @lc app=leetcode id=1140 lang=golang
 *
 * [1140] Stone Game II
 */

// @lc code=start
import (
	"fmt"
	"math"
)

type State struct {
	i int
	M int
}

var memo map[State]int

var gpiles []int

func stoneGameII(piles []int) int {
	gpiles = piles
	memo = make(map[State]int, 0)

	sum := 0
	for _, v := range piles {
		sum += v
	}
	return (sum + dp(0, 1)) / 2
}

// how much score the first hand exceed the second hand
func dp(i, M int) (mscore int) {
	state := State{i: i, M: M}
	if v, ok := memo[state]; ok {
		return v
	}

	// i + 2 * M - 1 >= len(gpiles) - 1
	if i+2*M >= len(gpiles) {
		mscore = 0
		for X := 1; X <= 2*M; X++ {
			if i+X-1 < len(gpiles) {
				mscore += gpiles[i+X-1]
			} else {
				break
			}
		}
		memo[state] = mscore
		return mscore
	}

	mscore = math.MinInt
	curscore := 0
	for X := 1; X <= 2*M; X++ {
		if i+X-1 < len(gpiles) {
			curscore += gpiles[i+X-1]
			mscore = max(mscore, curscore-dp(i+X, max(M, X)))
		} else {
			break
		}
	}

	memo[state] = mscore
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
