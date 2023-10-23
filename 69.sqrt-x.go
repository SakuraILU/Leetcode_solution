/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	left, right := 1, x/2

	for left <= right {
		mid := left + (right-left)/2
		n2 := mid * mid

		if x == n2 {
			return mid
		} else if x < n2 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}

// @lc code=end
