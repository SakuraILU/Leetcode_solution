/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start

import "fmt"

var res [][]int
var track []int
var cursum int
var gcandidates []int
var gtarget int

func combinationSum(candidates []int, target int) [][]int {
	gcandidates = candidates
	gtarget = target
	track = make([]int, 0)
	cursum = 0
	res = make([][]int, 0)

	dfs(0)

	return res
}

func dfs(start int) {
	if cursum > gtarget {
		return
	} else if cursum == gtarget {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		return
	}

	for i := start; i < len(gcandidates); i++ {
		track = append(track, gcandidates[i])
		cursum += gcandidates[i]
		dfs(i)
		cursum -= gcandidates[i]
		track = track[0 : len(track)-1]
	}
}

// @lc code=end
