/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start
func numTrees(n int) int {
	return backtrack(1, n)
}

func backtrack(l, h int) (num int) {
	if l >= h {
		return 1
	}

	for i := l; i <= h; i++ {
		lnum := backtrack(l, i-1)
		rnum := backtrack(i+1, h)
		num += lnum * rnum
	}

	return
}

// @lc code=end

