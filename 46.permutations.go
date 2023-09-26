/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
var gnums []int

var res [][]int
var track []int
var visited []bool

func permute(nums []int) [][]int {
	gnums = nums
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
		visited[i] = true
		track = append(track, gnums[i])

		dfs()

		track = track[0 : len(track)-1]
		visited[i] = false
	}
}

// @lc code=end
