/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	lastcolumn := make([]int, 0)
	rows, cols := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		lastcolumn = append(lastcolumn, matrix[i][cols-1])
	}
	row := bsearch(lastcolumn, target)
	if row >= len(lastcolumn) {
		return false
	}

	col := bsearch(matrix[row], target)
	return matrix[row][col] == target
}

func bsearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}
	}

	return left
}

// @lc code=end
