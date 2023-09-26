/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
var gnums []int

var res [][]int
var track []int

func subsetsWithDup(nums []int) [][]int {
	gnums = nums
	sort.Ints(gnums)
	res = make([][]int, 0)
	track = make([]int, 0)

	dfs(0)

	return res
}

func dfs(start int) {
	tmp := make([]int, len(track))
	copy(tmp, track)
	res = append(res, tmp)

	for i := start; i < len(gnums); i++ {
		if i > start && gnums[i] == gnums[i-1] {
			continue
		}

		track = append(track, gnums[i])

		dfs(i + 1)

		track = track[0 : len(track)-1]
	}
}

// @lc code=end
