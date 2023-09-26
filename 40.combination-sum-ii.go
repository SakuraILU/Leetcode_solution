/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
import (
	"fmt"
	"sort"
)

var res [][]int
var track []int
var cursum int
var gcandidates []int
var gtarget int

func combinationSum2(candidates []int, target int) [][]int {
	gcandidates = candidates
	sort.Ints(gcandidates)
	gtarget = target
	res = make([][]int, 0)
	track = make([]int, 0)
	cursum = 0

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
		if i > start && gcandidates[i] == gcandidates[i-1] {
			continue
		}

		track = append(track, gcandidates[i])
		cursum += gcandidates[i]

		dfs(i + 1)

		cursum -= gcandidates[i]
		track = track[0 : len(track)-1]
	}
}

// @lc code=end
