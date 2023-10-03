/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */

// @lc code=start
func reverseString(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp

		left++
		right--
	}
}

// @lc code=end

