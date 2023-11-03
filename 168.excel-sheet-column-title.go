/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start

func convertToTitle(columnNumber int) (res string) {
	M := 26

	for columnNumber > 0 {
		digit := columnNumber % M

		if digit > 0 {
			res = string(digit+'A'-1) + res
		} else {
			res = "Z" + res
			columnNumber -= M
		}

		columnNumber = columnNumber / M
	}

	return
}

// @lc code=end
