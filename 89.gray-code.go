/*
 * @lc app=leetcode id=89 lang=golang
 *
 * [89] Gray Code
 */

// @lc code=start
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	codes := grayCode(n - 1)

	length := len(codes)
	for i := length - 1; i >= 0; i-- {
		code := codes[i] | (1 << (n - 1))
		codes = append(codes, code)
	}

	return codes
}

// @lc code=end
