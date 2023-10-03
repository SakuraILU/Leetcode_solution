/*
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */

// @lc code=start
func hIndex(citations []int) int {
	left, right := 0, len(citations)
	for left <= right {
		mid := left + (right-left)/2

		cnt := 0
		for _, citation := range citations {
			if citation >= mid {
				cnt++
			}
		}

		if cnt == mid {
			return mid
		} else if cnt < mid {
			right = mid - 1
		} else if cnt > mid {
			left = mid + 1
		}
	}

	return right
}

// @lc code=end

