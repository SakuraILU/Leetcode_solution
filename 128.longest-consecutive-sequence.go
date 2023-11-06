/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
var memo map[int]int
var cnts map[int]int

func longestConsecutive(nums []int) int {
	memo = make(map[int]int)

	cnts = make(map[int]int)
	for _, num := range nums {
		cnts[num]++
	}

	mnum := 0
	for num, _ := range cnts {
		// get the number of consecutive elements
		mnum = max(mnum, getConsecutiveNumber(num))
	}

	return mnum
}

func getConsecutiveNumber(n int) int {
	if _, ok := cnts[n]; !ok {
		return 0
	}

	if v, ok := memo[n]; ok {
		return v
	}

	memo[n] = 1 + getConsecutiveNumber(n+1)
	return memo[n]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
