/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	need := make(map[rune]int, 0)
	for _, v := range s1 {
		need[v]++
	}

	cnts := make(map[rune]int, 0)
	l, r := 0, 0
	nvalid, nneed := 0, len(need)
	for ; r < len(s2); r++ {
		cnts[rune(s2[r])]++
		if _, ok := need[rune(s2[r])]; ok {
			if cnts[rune(s2[r])] == need[rune(s2[r])] {
				nvalid++
			}
		}

		if nvalid == nneed {
			return true
		}

		for (r - l + 1) >= len(s1) {
			if _, ok := need[rune(s2[l])]; ok {
				if cnts[rune(s2[l])] == need[rune(s2[l])] {
					nvalid--
				}
			}
			cnts[rune(s2[l])]--
			l++
		}
	}

	return false
}

// @lc code=end
