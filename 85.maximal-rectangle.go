/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */

// @lc code=start
var _matrix [][]byte

func maximalRectangle(matrix [][]byte) (marea int) {
	_matrix = matrix

	heights := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		updateHeight(heights, i)
		area := largestRectangleArea(heights)
		marea = max(marea, area)
	}

	return
}

func updateHeight(heights []int, row int) {
	for i := 0; i < len(heights); i++ {
		if _matrix[row][i] == '1' {
			heights[i]++
		} else {
			heights[i] = 0
		}
	}
}

func largestRectangleArea(heights []int) (marea int) {
	presmallers := findPreSmallers(heights)
	nextsmallers := findNextSmallers(heights)

	for i := 0; i < len(heights); i++ {
		area := heights[i] * (nextsmallers[i] - presmallers[i] - 1)
		marea = max(marea, area)
	}

	return
}

func findPreSmallers(heights []int) (smallers []int) {
	smallers = make([]int, len(heights))

	stack := make([]int, 0)
	for i := 0; i < len(heights); i++ {
		height := heights[i]
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= height {
			stack = stack[0 : len(stack)-1]
		}

		idx := -1
		if len(stack) > 0 {
			idx = stack[len(stack)-1]
		}

		smallers[i] = idx

		stack = append(stack, i)
	}

	return
}

func findNextSmallers(heights []int) (smallers []int) {
	smallers = make([]int, len(heights))

	stack := make([]int, 0)
	for i := len(heights) - 1; i >= 0; i-- {
		height := heights[i]
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= height {
			stack = stack[0 : len(stack)-1]
		}

		idx := len(heights)
		if len(stack) > 0 {
			idx = stack[len(stack)-1]
		}

		smallers[i] = idx

		stack = append(stack, i)
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
