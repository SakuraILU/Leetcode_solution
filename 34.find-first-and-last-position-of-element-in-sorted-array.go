/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	left := bsearch(nums, target, true)
	right := bsearch(nums, target, false)

	if left <= right {
		return []int{left, right}
	} else {
		return []int{-1, -1}
	}
}

func bsearch(nums []int, target int, leftboard bool) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if target == nums[mid] {
			if leftboard {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if leftboard {
		return left
	} else {
		return right
	}
}

// @lc code=end
