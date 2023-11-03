/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) (res string) {
	res = "1"
	for i := 1; i < n; i++ {
		res = dfs(res)
	}

	return
}

func dfs(s string) (res string) {
	c := s[0]
	i := 0
	for ; i < len(s) && s[i] == c; i++ {
	}

	res = strconv.Itoa(i) + string(c)
	if i < len(s) {
		res += dfs(s[i:len(s)])
	}
	return
}

// @lc code=end
