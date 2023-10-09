/*
 * @lc app=leetcode id=875 lang=golang
 *
 * [875] Koko Eating Bananas
 */

// @lc code=start
import "math"

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0
	for i := 0; i < len(piles); i++ {
		right = max(right, piles[i])
	}

	for left <= right {
		mid := left + (right-left)/2
		if canFinish(piles, mid, h) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func canFinish(piles []int, speed int, lefthours int) bool {
	for i := 0; i < len(piles); i++ {
		lefthours -= int(math.Ceil(float64(piles[i]) / float64(speed)))
		if lefthours < 0 {
			return false
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
