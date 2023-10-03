/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
	s = strings.Join(strings.Fields(s), " ")
	bytes := []byte(s)
	reverse(bytes)
	dfs(bytes)
	return string(bytes)
}

func reverse(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp

		left++
		right--
	}
}

func dfs(s []byte) {
	right := 0
	for right < len(s) && s[right] != ' ' {
		right++
	}

	next := right + 1

	reverse(s[0:right])

	if next < len(s) {
		dfs(s[next:len(s)])
	}
}

// @lc code=end
