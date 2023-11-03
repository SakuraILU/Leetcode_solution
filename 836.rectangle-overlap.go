/*
 * @lc app=leetcode id=836 lang=golang
 *
 * [836] Rectangle Overlap
 */

// @lc code=start

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return isLineOverlap([]int{rec1[0], rec1[2]}, []int{rec2[0], rec2[2]}) && isLineOverlap([]int{rec1[1], rec1[3]}, []int{rec2[1], rec2[3]})
}

func isLineOverlap(line1 []int, line2 []int) bool {
	return !(line1[1] <= line2[0] || line1[0] >= line2[1])
}

// @lc code=end
