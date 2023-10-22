/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	for i := left; i <= right; i++ {
		if nums[i] == 0 {
			tmp := nums[i]
			nums[i] = nums[left]
			nums[left] = tmp
			left++
		} else if nums[i] == 2 {
			tmp := nums[i]
			nums[i] = nums[right]
			nums[right] = tmp
			right--

			// swap a new value i not traversed yet, check again
			i--
		}
	}

	return
}

// @lc code=end

