/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start

var _matrix [][]int

func searchMatrix(matrix [][]int, target int) bool {
	_matrix = matrix

	return bsearch(0, len(matrix[0])-1, target)
}

func bsearch(row int, cols int, target int) bool {
	if row >= len(_matrix) || cols < 0 {
		return false
	}

	left, right := 0, cols
	for left <= right {
		mid := left + (right-left)/2
		if _matrix[row][mid] == target {
			return true
		} else if _matrix[row][mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return bsearch(row+1, right, target)
}

// @lc code=end
