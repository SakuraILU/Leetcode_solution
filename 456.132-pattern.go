/*
 * @lc app=leetcode id=456 lang=golang
 *
 * [456] 132 Pattern
 */

// @lc code=start
func find132pattern(nums []int) bool {
	queue := make([]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		for len(queue) > 0 && nums[i] <= queue[len(queue)-1] {
			queue = queue[0 : len(queue)-1]
		}
		queue = append(queue, nums[i])
		low := queue[0]
		top := nums[i+1]

		if top <= low {
			continue
		}

		for j := i + 2; j < len(nums); j++ {
			if nums[j] > low && nums[j] < top {
				return true
			}
		}
	}

	return false
}

// @lc code=end
