/*
 * @lc app=leetcode id=1011 lang=golang
 *
 * [1011] Capacity To Ship Packages Within D Days
 */

// @lc code=start

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	for _, weight := range weights {
		left = max(left, weight)
		right += weight
	}

	for left <= right {
		mid := left + (right-left)/2

		if canShipWithinDays(mid, weights, days) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func canShipWithinDays(capacity int, weights []int, leftdays int) bool {
	volume := capacity
	for _, weight := range weights {
		if volume >= weight {
			volume -= weight
		} else {
			leftdays--
			volume = capacity

			if leftdays <= 0 {
				return false
			}
			volume -= weight
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
