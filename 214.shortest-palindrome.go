/*
 * @lc app=leetcode id=214 lang=golang
 *
 * [214] Shortest Palindrome
 */

// @lc code=start
var _s string

func shortestPalindrome(s string) string {
	_s = s

	len2match := len(s) - dp(0)
	return reverse(s[len(s)-len2match:len(s)]) + s
}

func reverse(s string) (rs string) {
	for i := len(s) - 1; i >= 0; i-- {
		rs += string(s[i])
	}

	return
}

func dp(i int) int {
	if i == len(_s) {
		return 0
	}

	j := i + dp(i+1) + 1
	if j < len(_s) && _s[i] == _s[j] {
		return j - i + 1
	}

	for k := min(j, len(_s)-1); k >= i; k-- {
		if isPalindrome(_s[i : k+1]) {
			return k - i + 1
		}
	}

	return 1
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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
