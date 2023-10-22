/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
// package leetcodesolution

import (
	"strconv"
	"strings"
)

var _s string
var track []string
var res []string

func restoreIpAddresses(s string) []string {
	_s = s
	track = make([]string, 0)
	res = make([]string, 0)

	dfs(0)

	return res
}

func dfs(start int) {
	if len(track) == 4 {
		if start >= len(_s) {
			res = append(res, strings.Join(track, "."))
		}
		return
	}

	for i := start; i < min(len(_s), start+3); i++ {
		subnet := _s[start : i+1]
		if !isValidSubnet(subnet) {
			continue
		}

		track = append(track, subnet)

		dfs(i + 1)

		track = track[0 : len(track)-1]
	}
}

func isValidSubnet(subnet string) bool {
	// can not be something like 000, 001, 01, 010...
	// that is to say, if subnet start with 0, it can only be "0"
	if strings.HasPrefix(subnet, "0") {
		return subnet == "0"
	}

	if n, _ := strconv.Atoi(subnet); n > 0 && n <= 255 {
		return true
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
