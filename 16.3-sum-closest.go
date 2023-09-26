/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
// package leetcodesolution

import "math"

var _nums []int
var _target int
var cursum int
var depth int

var res int

func threeSumClosest(nums []int, target int) int {
	_nums = nums
	_target = target
	cursum = 0
	depth = 0
	res = math.MaxInt

	dfs(0)

	return res
}

func dfs(start int) {
	if depth == 3 {
		if Abs(cursum-_target) < Abs(res-_target) {
			res = cursum
		}
		return
	}

	for i := start; i < len(_nums); i++ {
		cursum += _nums[i]
		depth++

		dfs(i + 1)

		depth--
		cursum -= _nums[i]
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
