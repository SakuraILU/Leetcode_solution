/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if target == nums[mid] {
			return mid
		} else {
			if nums[mid] > nums[right] {
				// left, mid, pivot, right
				if target >= nums[left] && target < nums[mid] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else {
				// left, pivot, mid, right
				// or left, right
				if target > nums[mid] && target <= nums[right] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}

	return -1
}

// @lc code=end

