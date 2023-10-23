/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {
	res := 0

	for _, num := range nums {
		res = res ^ num
	}

	return res
}

// @lc code=end

