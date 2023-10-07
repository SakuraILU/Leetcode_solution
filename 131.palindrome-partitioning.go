/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
var _s string
var track []string
var res [][]string

func partition(s string) [][]string {
	_s = s
	track = make([]string, 0)
	res = make([][]string, 0)

	dfs(0)

	return res
}

func dfs(start int) {
	if start >= len(_s) {
		tmp := make([]string, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		return
	}

	for i := start; i < len(_s); i++ {
		substr := _s[start : i+1]
		if !isPalindrome(substr) {
			continue
		}

		track = append(track, substr)
		dfs(i + 1)
		track = track[0 : len(track)-1]
	}
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// @lc code=end
