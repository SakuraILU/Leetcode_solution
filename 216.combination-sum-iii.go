/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

// @lc code=start

var nums []int
var target int
var gk int

var res [][]int
var track []int
var cursum int

func combinationSum3(k int, n int) [][]int {
	target = n
	gk = k
	nums = make([]int, 0)
	for i := 1; i < 10; i++ {
		nums = append(nums, i)
	}

	res = make([][]int, 0)
	track = make([]int, 0)
	cursum = 0

	dfs(0)

	return res
}

func dfs(start int) {
	if cursum > target || len(track) > gk {
		return
	} else if cursum == target {
		if len(track) == gk {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
		}
		return
	}

	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		cursum += nums[i]

		dfs(i + 1)

		cursum -= nums[i]
		track = track[0 : len(track)-1]

	}

	return
}

// @lc code=end
