/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1] {
			return true
		}

		return false
	})

	return merge(intervals)
}

func merge(intervals [][]int) (merged [][]int) {
	if len(intervals) == 1 {
		merged = make([][]int, 0)
		merged = append(merged, intervals...)
		return
	}

	merged = merge(intervals[0 : len(intervals)-1])

	cur := intervals[len(intervals)-1]
	pre := merged[len(merged)-1]

	if cur[1] <= pre[1] {
		return
	} else if cur[0] <= pre[1] {
		merged[len(merged)-1][1] = cur[1]
	} else {
		merged = append(merged, cur)
	}

	return
}

// @lc code=end
