/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start

func decodeString(s string) string {
	return dfs(s)
}

func dfs(s string) (res string) {
	if len(s) < 3 {
		return s
	}

	if !unicode.IsDigit(rune(s[0])) {
		return s[0:1] + dfs(s[1:])
	}

	i, j := 0, 0
	for ; i < len(s) && s[i] != '['; i++ {
	}

	k, _ := strconv.Atoi(s[0:i])

	depth := 0
	for j = i; j < len(s); j++ {
		if s[j] == '[' {
			depth++
		} else if s[j] == ']' {
			depth--

			if depth == 0 {
				break
			}
		}
	}

	substr := dfs(s[i+1 : j])

	for i := 0; i < k; i++ {
		res += substr
	}

	if j+1 < len(s) {
		res += dfs(s[j+1 : len(s)])
	}

	return res
}

// @lc code=end
