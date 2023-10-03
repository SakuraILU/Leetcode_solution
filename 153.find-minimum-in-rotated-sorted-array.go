/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if mid > 0 && nums[mid-1] > nums[mid] {
			return nums[mid]
		} else {
			if nums[mid] > nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return nums[left]
}

// @lc code=end

