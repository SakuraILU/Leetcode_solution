/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */

// @lc code=start
func reverseWords(s string) string {
	bytes := []byte(s)

	dfs(bytes)

	return string(bytes)
}

func dfs(s []byte) {
	left, right := 0, 0
	for right < len(s) && s[right] != ' ' {
		right++
	}

	next_word := right + 1

	right--
	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp

		left++
		right--
	}

	if next_word >= len(s) {
		return
	}
	dfs(s[next_word:])
}

// @lc code=end
