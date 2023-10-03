/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start
func findDuplicate(nums []int) int {
	left, right := 1, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		lesscnt := 0
		equalcnt := 0
		for _, num := range nums {
			if num == mid {
				equalcnt++
			} else if num < mid {
				lesscnt++
			}
		}

		if equalcnt > 1 {
			return mid
		} else if lesscnt >= mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

// @lc code=end

