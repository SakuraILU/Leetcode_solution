/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	left, right := 0, 0
	for _, num := range nums {
		left = min(left, num)
		right = max(right, num)
	}

	for left <= right {
		mid := left + (right-left)/2

		lesscnt := 0
		equalcnt := 0
		for _, num := range nums {
			if num < mid {
				lesscnt++
			} else if num == mid {
				equalcnt++
			}
		}

		if equalcnt > len(nums)/2 {
			return mid
		} else if lesscnt > len(nums)/2 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
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

