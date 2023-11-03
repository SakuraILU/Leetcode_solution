/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
// package leetcodesolution

// import (
//
//	"strconv"
//	"strings"
//
// )

var memo map[int]int

var _s string

func numDecodings(s string) (cnt int) {
	_s = s

	memo = make(map[int]int)
	return dp(0)
}

func dp(start int) (cnt int) {
	if start >= len(_s) {
		return 1
	}

	if v, ok := memo[start]; ok {
		return v
	}

	for i := start; i < min(start+2, len(_s)); i++ {
		snum := _s[start : i+1]
		if strings.HasPrefix(snum, "0") {
			continue
		}
		num, _ := strconv.Atoi(_s[start : i+1])
		if num <= 0 || num > 26 {
			continue
		}

		cnt += dp(i + 1)
	}

	memo[start] = cnt
	return
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
