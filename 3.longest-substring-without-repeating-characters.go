/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start

func lengthOfLongestSubstring(s string) int {
	cnts := make(map[byte]int)
	nduplica := 0
	mlen := 0

	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		if cnts[c] == 1 {
			nduplica++
		}
		cnts[c]++

		right++
		for nduplica > 0 {
			c := s[left]
			cnts[c]--
			if cnts[c] == 1 {
				nduplica--
			}

			left++
		}

		if nduplica == 0 {
			mlen = max(mlen, right-left)
		}
	}

	return mlen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
