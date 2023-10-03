/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1] {
			return true
		}
		return false
	})

	merged := dfs(intervals)
	// fmt.Println(merged)
	return merged
}

func dfs(intervals [][]int) (merged [][]int) {
	if len(intervals) == 1 {
		merged = make([][]int, 0)
		merged = append(merged, intervals[0])
		return
	}

	merged = dfs(intervals[0 : len(intervals)-1])
	cur := intervals[len(intervals)-1]
	pre := merged[len(merged)-1]
	// fmt.Println(pre, cur)

	// if cur[1] < pre[1] do nothing..
	if cur[1] < pre[1] {
	} else if cur[0] <= pre[1] {
		merged[len(merged)-1][1] = cur[1]
	} else {
		merged = append(merged, cur)
	}

	return
}

// @lc code=end

