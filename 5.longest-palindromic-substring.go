/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
var memo map[int]int

var _s string

func longestPalindrome(s string) string {
	_s = s

	memo = make(map[int]int, 0)
	mlen := 0
	res := ""
	for i := 0; i < len(s); i++ {
		len := dp(i)
		if len > mlen {
			mlen = len
			res = s[i-len+1 : i+1]
		}
	}

	return res
}

func dp(j int) int {
	if j == 0 {
		return 1
	}

	if v, ok := memo[j]; ok {
		return v
	}

	mlen := 1
	i := j - dp(j-1) - 1
	for k := max(0, i); k < j; k++ {
		if _s[k] == _s[j] && isPalindrome(k, j) {
			return j - k + 1
		}
	}

	memo[j] = mlen
	return mlen
}

func isPalindrome(i, j int) bool {
	for i < j {
		if _s[i] != _s[j] {
			return false
		}
		i++
		j--
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

// @lc code=end
