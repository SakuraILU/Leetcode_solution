/*
 * @lc app=leetcode id=140 lang=golang
 *
 * [140] Word Break II
 */

// @lc code=start
var _s string
var _wordDIc []string
var track []string
var res []string

func wordBreak(s string, wordDict []string) []string {
	_s = s
	_wordDIc = wordDict

	track = make([]string, 0)
	res = make([]string, 0)

	dfs(0)

	return res
}

func dfs(start int) {
	if start >= len(_s) {
		res = append(res, strings.Join(track, " "))
		return
	}

	for _, word := range _wordDIc {
		if !strings.HasPrefix(_s[start:len(_s)], word) {
			continue
		}

		track = append(track, word)

		dfs(start + len(word))

		track = track[0 : len(track)-1]
	}

	return
}

// @lc code=end
