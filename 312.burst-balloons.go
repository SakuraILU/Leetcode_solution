/*
 * @lc app=leetcode id=312 lang=golang
 *
 * [312] Burst Balloons
 */

// @lc code=start
type State struct {
	i, j int
}

var memo map[State]int

var gnums []int

func maxCoins(nums []int) int {
	gnums = make([]int, len(nums)+2)
	copy(gnums[1:len(nums)+1], nums)
	gnums[0], gnums[len(gnums)-1] = 1, 1

	memo = make(map[State]int, 0)
	return dp(1, len(nums))
}

func dp(i, j int) (mscore int) {
	if i > j {
		return 0
	} else if i == j {
		return gnums[i-1] * gnums[i] * gnums[i+1]
	}

	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	mscore = 0
	for k := i; k <= j; k++ {
		mscore = max(mscore, dp(i, k-1)+dp(k+1, j)+gnums[i-1]*gnums[k]*gnums[j+1])
	}

	memo[state] = mscore
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
