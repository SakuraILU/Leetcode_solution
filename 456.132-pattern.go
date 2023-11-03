/*
 * @lc app=leetcode id=456 lang=golang
 *
 * [456] 132 Pattern
 */

// @lc code=start

func find132pattern(nums []int) bool {
	leftmins := make([]int, len(nums))
	leftmins[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		leftmins[i] = min(leftmins[i-1], nums[i])
	}

	decstack := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for len(decstack) > 0 && nums[i] >= nums[decstack[len(decstack)-1]] {
			decstack = decstack[0 : len(decstack)-1]
		}

		if len(decstack) > 0 {
			second := decstack[len(decstack)-1]
			if second > 0 && leftmins[second-1] < nums[i] {
				return true
			}
		}

		decstack = append(decstack, i)
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
