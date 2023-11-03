/*
 * @lc app=leetcode id=354 lang=golang
 *
 * [354] Russian Doll Envelopes
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return false
	})

	dp := make([]int, len(envelopes))

	mlen := 0
	for i, envelope := range envelopes {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelope[0] > envelopes[j][0] && envelope[1] > envelopes[j][1] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}

		mlen = max(mlen, dp[i])
	}

	return mlen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
