/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] Zigzag Conversion
 */

// @lc code=start
// interval1 interval2
// (n-1) * 2 0
// (n-2) * 2 1 * 2
// (n-3) * 2 2 * 2
// ...
// 0	  *2  (n-1) * 2

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := ""
	for i := 0; i < numRows; i++ {
		pos := i
		cnt := 0
		for pos < len(s) {
			res += string(s[pos])

			if i == 0 {
				pos += (numRows - i - 1) * 2
			} else if i == numRows-1 {
				pos += i * 2
			} else if cnt%2 == 0 {
				pos += (numRows - i - 1) * 2
			} else {
				pos += i * 2
			}
			cnt++
		}
	}

	return res
}

// @lc code=end
