/*
 * @lc app=leetcode id=395 lang=golang
 *
 * [395] Longest Substring with At Least K Repeating Characters
 */

// @lc code=start
// package leetcodesolution

import (
	"strings"
)

func longestSubstring(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	cnts := make(map[rune]int, 0)
	for _, c := range s {
		cnts[c]++
	}

	var cut_c rune
	for c, cnt := range cnts {
		if cnt < k {
			cut_c = c
			break
		}
	}

	cut_p := strings.Index(s, string(cut_c))
	if cut_p == -1 {
		return len(s)
	}

	// [0, cut_p)
	left_max := longestSubstring(s[0:cut_p], k)
	// (cut_p, end)
	right_max := longestSubstring(s[cut_p+1:], k)

	return max(left_max, right_max)
}

func max(a, b int) (res int) {
	if a > b {
		res = a
	} else {
		res = b
	}
	return
}

// @lc code=end
