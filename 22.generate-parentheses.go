/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start

var gn int

var choices []byte = []byte{'(', ')'}
var res []string
var track []byte

func generateParenthesis(n int) []string {
	gn = n
	res = make([]string, 0)
	track = make([]byte, 0)

	dfs(0, 0)

	return res
}

func dfs(left, right int) {
	if left < right || left > gn || right > gn {
		return
	} else if left == gn && right == gn {
		tmp := string(track)
		res = append(res, tmp)
		return
	}

	for _, choice := range choices {
		track = append(track, choice)

		switch choice {
		case '(':
			dfs(left+1, right)
		case ')':
			dfs(left, right+1)
		}

		track = track[0 : len(track)-1]
	}
}

// @lc code=end
