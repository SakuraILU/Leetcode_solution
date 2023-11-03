/*
 * @lc app=leetcode id=164 lang=golang
 *
 * [164] Maximum Gap
 */

// @lc code=start
// package leetcodesolution

import "math"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	minval, maxval := math.MaxInt, math.MinInt
	for _, num := range nums {
		minval = min(minval, num)
		maxval = max(maxval, num)
	}

	mins := make([]int, len(nums)+1)
	maxs := make([]int, len(nums)+1)
	for i := 0; i < len(mins); i++ {
		mins[i] = math.MaxInt
		maxs[i] = math.MinInt
	}
	// n+1 buckets, store num from min to max
	lenbucket := int(math.Ceil(float64(maxval-minval+1) / float64(len(nums)+1)))
	for _, num := range nums {
		idx := (num - minval) / lenbucket
		mins[idx] = min(mins[idx], num)
		maxs[idx] = max(maxs[idx], num)
	}

	maxgap := 0
	prebucket_idx := -1
	for i := 0; i < len(mins); i++ {
		if mins[i] == math.MaxInt {
			continue
		}

		if prebucket_idx >= 0 {
			maxgap = max(maxgap, mins[i]-maxs[prebucket_idx])
		}

		prebucket_idx = i
	}

	return maxgap
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
