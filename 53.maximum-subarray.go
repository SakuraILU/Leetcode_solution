/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
var memo map[int]int

var _nums []int

func maxSubArray(nums []int) int {
	_nums = nums
	memo = make(map[int]int, 0)

	msum := math.MinInt
	for i := 0; i < len(_nums); i++ {
		msum = max(msum, dp(i))
	}

	return msum
}

func dp(i int) (msum int) {
	if i < 0 {
		return 0
	} else if i == 0 {
		return _nums[i]
	}

	if v, ok := memo[i]; ok {
		return v
	}

	msum = max(dp(i-1)+_nums[i], _nums[i])

	memo[i] = msum
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
