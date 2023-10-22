/*
 * @lc app=leetcode id=87 lang=golang
 *
 * [87] Scramble String
 */

// @lc code=start

var memo map[string]bool

func isScramble(s1 string, s2 string) bool {
	memo = make(map[string]bool, 0)

	return dfs(s1, s2)
}

func dfs(s1, s2 string) bool {
	state := s1 + "," + s2
	if v, ok := memo[state]; ok {
		return v
	}

	if s1 == s2 {
		memo[state] = true
		return true
	}

	for i := 1; i < len(s1); i++ {
		// [0, i) [i, len(s1))
		sub1 := s1[0:i]
		sub2 := s1[i:len(s1)]

		if dfs(sub1, s2[0:i]) && dfs(sub2, s2[i:len(s1)]) {
			memo[state] = true
			return true
		} else if dfs(sub1, s2[len(s1)-i:len(s1)]) && dfs(sub2, s2[0:len(s1)-i]) {
			memo[state] = true
			return true
		}
	}

	memo[state] = false
	return false
}

// @lc code=end
