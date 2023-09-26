/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
var gnums []int

var res [][]int
var track []int

func subsets(nums []int) [][]int {
	gnums = nums
	track = make([]int, 0)
	res = make([][]int, 0)

	dfs(0)

	return res
}

func dfs(start int) {
	tmp := make([]int, len(track))
	copy(tmp, track)
	res = append(res, tmp)

	for i := start; i < len(gnums); i++ {
		track = append(track, gnums[i])

		dfs(i + 1)

		track = track[0 : len(track)-1]
	}
}

// @lc code=end
