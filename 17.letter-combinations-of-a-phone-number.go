/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start

var table map[byte][]byte = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

var _digits string
var track []byte
var res []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	_digits = digits
	track = make([]byte, 0)
	res = make([]string, 0)

	dfs(0)

	return res
}

func dfs(i int) {
	if i >= len(_digits) {
		res = append(res, string(track))
		return
	}

	choices := table[_digits[i]]
	for _, c := range choices {
		track = append(track, c)

		dfs(i + 1)

		track = track[0 : len(track)-1]
	}
}

// @lc code=end
