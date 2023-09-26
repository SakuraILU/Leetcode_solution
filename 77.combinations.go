/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start

var track []int
var res [][]int
var gn, gk int

func combine(n int, k int) [][]int {
	gn, gk = n, k
	track = make([]int, 0)
	res = make([][]int, 0)

	dfs(1)

	return res
}

func dfs(start int) {
	if len(track) == gk {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		return
	}

	for i := start; i <= gn; i++ {
		track = append(track, i)
		dfs(i + 1)
		track = track[0 : len(track)-1]
	}
}

// @lc code=end
