/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
func findPeakElement(nums []int) int {
	nums = append([]int{math.MinInt}, append(nums, math.MinInt)...)
	left, right := 1, len(nums)-2
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid - 1
		} else if nums[mid] < nums[mid-1] {
			right = mid - 1
		} else if nums[mid] < nums[mid+1] {
			left = mid + 1
		}
	}

	return -1
}

// @lc code=end
