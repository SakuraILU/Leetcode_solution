/*
 * @lc app=leetcode id=1288 lang=golang
 *
 * [1288] Remove Covered Intervals
 */

// @lc code=start
// package leetcodesolution

import (
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1] {
			return true
		}
		return false
	})

	remain := dfs(intervals)
	return len(remain)
}

func dfs(intervals [][]int) (remain [][]int) {
	if len(intervals) == 1 {
		remain = make([][]int, 0)
		remain = append(remain, intervals[0])
		return
	}

	remain = dfs(intervals[0 : len(intervals)-1])

	cur := intervals[len(intervals)-1]
	pre := remain[len(remain)-1]

	if cur[1] <= pre[1] {
		return
	} else if cur[0] == pre[0] {
		remain[len(remain)-1] = cur
	} else {
		remain = append(remain, cur)
	}

	return
}

// @lc code=end
