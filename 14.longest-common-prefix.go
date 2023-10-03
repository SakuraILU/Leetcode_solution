/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	return dfs(strs)
}

func dfs(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	str := strs[0]
	prefix := dfs(strs[1:])

	curpos := 0
	for curpos < len(strs[0]) && curpos < len(prefix) {
		if str[curpos] == prefix[curpos] {
			curpos++
		} else {
			break
		}
	}

	return str[0:curpos]
}

// @lc code=end
