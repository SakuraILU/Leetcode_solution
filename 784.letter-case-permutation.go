/*
 * @lc app=leetcode id=784 lang=golang
 *
 * [784] Letter Case Permutation
 */

// @lc code=start
var _s string

var track []byte
var res []string

func letterCasePermutation(s string) []string {
	_s = s

	track = make([]byte, 0)
	res = make([]string, 0)

	dfs(0)

	return res
}

func dfs(i int) {
	if i == len(_s) {
		tmp := string(track)
		res = append(res, tmp)
		return
	}

	track = append(track, _s[i])
	dfs(i + 1)
	track = track[0 : len(track)-1]

	if unicode.IsUpper(rune(_s[i])) {
		track = append(track, byte(unicode.ToLower(rune(_s[i]))))
		dfs(i + 1)
		track = track[0 : len(track)-1]
	} else if unicode.IsLower(rune(_s[i])) {
		track = append(track, byte(unicode.ToUpper(rune(_s[i]))))
		dfs(i + 1)
		track = track[0 : len(track)-1]
	}
}

// @lc code=end
