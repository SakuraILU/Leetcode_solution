/*
 * @lc app=leetcode id=424 lang=golang
 *
 * [424] Longest Repeating Character Replacement
 */

// @lc code=start
func characterReplacement(s string, k int) int {
	l, r := 0, 0

	longest_len := 0
	max_freq := 0
	cnts := make(map[rune]int, 0)
	for ; r < len(s); r++ {
		cnts[rune(s[r])]++

		if cnts[rune(s[r])] > max_freq {
			max_freq = cnts[rune(s[r])]
		}

		cur_len := r - l + 1
		if cur_len-max_freq > k {
			cnts[rune(s[l])]--
			l++
		} else {
			if cur_len > longest_len {
				longest_len = cur_len
			}
		}
	}

	return longest_len
}

// @lc code=end

