/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	storei, storej := -1, -1

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				storei, storej = i, j

				for k := i + 1; k < len(matrix); k++ {
					if matrix[k][j] != 0 {
						matrix[k][j] = 0
					} else {
						matrix[k][j] = 1
					}
				}

				for k := 0; k < len(matrix[0]); k++ {
					if matrix[i][k] != 0 {
						matrix[i][k] = 0
					} else {
						matrix[i][k] = 1
					}
				}
				goto outer
			}
		}
	}
outer:

	if storei == -1 {
		return
	}

	for i := storei + 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j != storej && matrix[i][j] == 0 {
				matrix[i][storej] = 1
				matrix[storei][j] = 1
			}
		}
	}

	for i := storei + 1; i < len(matrix); i++ {
		if matrix[i][storej] == 1 {
			clearRow(matrix, i)
		}
	}

	for j := 0; j < len(matrix[0]); j++ {
		if matrix[storei][j] == 1 {
			clearCol(matrix, j)
		}
	}

	clearRow(matrix, storei)
	clearCol(matrix, storej)

	return
}

func clearCol(matrix [][]int, j int) {
	for k := 0; k < len(matrix); k++ {
		matrix[k][j] = 0
	}
}

func clearRow(matrix [][]int, i int) {
	for k := 0; k < len(matrix[0]); k++ {
		matrix[i][k] = 0
	}
}

// @lc code=end
