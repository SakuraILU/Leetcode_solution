/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
// package leetcodesolution

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)

	need := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	length := len(p)
	cnts := make(map[byte]int)
	nvalid := 0
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		if v, ok := need[c]; ok {
			cnts[c]++
			if cnts[c] == v {
				nvalid++
			}
		}

		right++

		if right-left == length {
			if nvalid == len(need) {
				res = append(res, left)
			}

			c := s[left]
			if v, ok := need[c]; ok {
				if cnts[c] == v {
					nvalid--
				}
				cnts[c]--
			}

			left++
		}
	}

	return res
}

// @lc code=end
