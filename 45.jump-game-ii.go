/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
var gnums []int

const INVALIDE = -1

var memo map[int]int

func jump(nums []int) int {
	gnums = nums
	memo = make(map[int]int, 0)

	return dp(0)
}

func dp(pos int) int {
	if pos == len(gnums)-1 {
		return 0
	}
	if pos+gnums[pos] >= len(gnums)-1 {
		return 1
	}

	if v, ok := memo[pos]; ok {
		return v
	}
	memo[pos] = INVALIDE

	mstep := math.MaxInt
	for jumpstep := 1; jumpstep <= gnums[pos]; jumpstep++ {
		if dp(pos+jumpstep) != INVALIDE {
			mstep = min(mstep, 1+dp(pos+jumpstep))
		}
	}

	if mstep == math.MaxInt {
		return INVALIDE
	}

	memo[pos] = mstep
	return mstep
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end

