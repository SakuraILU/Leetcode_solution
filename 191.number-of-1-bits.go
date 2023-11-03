/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {
	i := 0
	for ; i < 32 && num != 0; i++ {
		num = num & (num - 1)
	}

	return i
}

// @lc code=end
