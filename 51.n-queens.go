/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start

var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)

	var rowtemp string
	for i := 0; i < n; i++ {
		rowtemp += "."
	}

	table := make([]string, n)
	for i := 0; i < len(table); i++ {
		table[i] = rowtemp
	}

	dfs(0, table)

	return res
}

func dfs(row int, table []string) {
	fmt.Println(table)
	if row >= len(table) {
		dsttable := make([]string, len(table))
		copy(dsttable, table)
		res = append(res, dsttable)
		return
	}

	for col := 0; col < len(table[row]); col++ {
		tablerow := []byte(table[row])
		tablerow[col] = 'Q'
		table[row] = string(tablerow)

		if isValid(row, col, table) {
			dfs(row+1, table)
		}

		tablerow = []byte(table[row])
		tablerow[col] = '.'
		table[row] = string(tablerow)
	}
}

func isValid(row, col int, table []string) bool {
	for i := 1; i <= row; i++ {
		if table[row-i][col] == 'Q' {
			return false
		}

		if col-i >= 0 && table[row-i][col-i] == 'Q' {
			return false
		}

		if col+i < len(table[0]) && table[row-i][col+i] == 'Q' {
			return false
		}
	}
	return true
}

// @lc code=end
