/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	set := make(map[int]interface{})

	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		} else {
			set[num] = struct{}{}
		}
	}

	return false
}

// @lc code=end

