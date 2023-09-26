/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start

import "sort"

var gnums []int
var res [][]int
var track []int
var visited []bool

func permuteUnique(nums []int) [][]int {
	gnums = nums
	sort.Ints(gnums)
	res = make([][]int, 0)
	track = make([]int, 0)
	visited = make([]bool, len(gnums))

	dfs()

	return res
}

func dfs() {
	if len(track) == len(gnums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		return
	}

	for i := 0; i < len(gnums); i++ {
		if visited[i] {
			continue
		}
		if i > 0 && gnums[i] == gnums[i-1] && !visited[i-1] {
			continue
		}

		visited[i] = true
		track = append(track, gnums[i])

		dfs()

		track = track[0 : len(track)-1]
		visited[i] = false
	}
}

// @lc code=end
