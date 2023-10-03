/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
var memo map[int]int

var _nums []int

func lengthOfLIS(nums []int) int {
	_nums = nums

	memo = make(map[int]int, 0)
	mlen := 0
	for i := 0; i < len(nums); i++ {
		mlen = max(mlen, dp(i))
	}

	return mlen
}

func dp(i int) (mlen int) {
	if v, ok := memo[i]; ok {
		return v
	}

	for k := i + 1; k < len(_nums); k++ {
		if _nums[k] > _nums[i] {
			mlen = max(mlen, dp(k))
		}
	}
	mlen += 1

	memo[i] = mlen
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
