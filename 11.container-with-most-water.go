/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height)-1

	most_water := 0
	for left < right {
		water := (right - left) * min(height[left], height[right])
		most_water = max(most_water, water)

		if height[left] <= height[right] {
			left++
		} else if height[left] > height[right] {
			right--
		}
	}

	return most_water
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end

