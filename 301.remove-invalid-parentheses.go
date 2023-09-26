/*
 * @lc app=leetcode id=301 lang=golang
 *
 * [301] Remove Invalid Parentheses
 */

// @lc code=start
// package leetcodesolution

var gs string
var res []string
var track []byte

func removeInvalidParentheses(s string) []string {
	gs = s
	res = make([]string, 0)
	track = make([]byte, 0)

	dfs(0, 0)

	return res
}

func dfs(start int, depth int) {
	if depth < 0 {
		return
	} else if depth == 0 {
		if len(res) == 0 || len(track) == len(res[0]) {
			res = append(res, string(track))
		} else if len(track) > len(res[0]) {
			res = make([]string, 0)
			res = append(res, string(track))
		}
	}

	for i := start; i < len(gs); i++ {
		if i > start && gs[i] == gs[i-1] {
			continue
		}

		track = append(track, gs[i])
		switch gs[i] {
		case '(':
			dfs(i+1, depth+1)
		case ')':
			dfs(i+1, depth-1)
		default:
			dfs(i+1, depth)
		}
		track = track[0 : len(track)-1]
	}
}

// @lc code=end
