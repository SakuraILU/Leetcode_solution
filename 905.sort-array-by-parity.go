/*
 * @lc app=leetcode id=905 lang=golang
 *
 * [905] Sort Array By Parity
 */

// @lc code=start
func sortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		for left < len(nums) && nums[left]%2 == 0 {
			left++
		}

		for right >= 0 && nums[right]%2 != 0 {
			right--
		}

		if left >= right {
			break
		}

		leftval := nums[left]
		nums[left] = nums[right]
		nums[right] = leftval
	}

	return nums
}

// @lc code=end

