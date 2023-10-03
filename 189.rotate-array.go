/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int) {
	k = k % len(nums)

	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:len(nums)])
}

func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		tmp := nums[left]
		nums[left] = nums[right]
		nums[right] = tmp

		left++
		right--
	}
}

// @lc code=end
