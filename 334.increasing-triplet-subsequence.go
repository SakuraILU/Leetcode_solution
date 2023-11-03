/*
 * @lc app=leetcode id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	leftmins := make([]int, len(nums))
	leftmins[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		leftmins[i] = min(leftmins[i-1], nums[i])
	}

	rightmaxs := make([]int, len(nums))
	rightmaxs[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		rightmaxs[i] = max(rightmaxs[i+1], nums[i])
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > leftmins[i-1] && nums[i] < rightmaxs[i+1] {
			return true
		}
	}

	return false
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
