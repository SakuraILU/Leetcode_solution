/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	return dfs(num, 32)
}

func dfs(num uint32, n int) uint32 {
	if n == 1 {
		return num
	}

	subn := n / 2
	lnum := num & (1<<subn - 1)
	rnum := num >> subn

	return dfs(lnum, subn)<<subn + dfs(rnum, subn)
}

// @lc code=end

