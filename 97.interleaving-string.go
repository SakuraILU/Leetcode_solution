/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 */

// @lc code=start
type State struct {
	i, j, k int
}

var memo map[State]bool

var _s1, _s2, _s3 string

func isInterleave(s1 string, s2 string, s3 string) bool {
	_s1, _s2, _s3 = s1, s2, s3

	memo = make(map[State]bool, 0)
	return dp(len(_s1)-1, len(_s2)-1, len(_s3)-1)
}

func dp(i, j, k int) (isInterleaving bool) {
	state := State{i: i, j: j, k: k}
	if v, ok := memo[state]; ok {
		return v
	}

	if i < 0 {
		isInterleaving = _s2[0:j+1] == _s3[0:k+1]
		memo[state] = isInterleaving
		return
	} else if j < 0 {
		isInterleaving = _s1[0:i+1] == _s3[0:k+1]
		memo[state] = isInterleaving
		return
	} else if k < 0 {
		isInterleaving = false
		memo[state] = isInterleaving
		return
	}

	if _s1[i] == _s3[k] && dp(i-1, j, k-1) {
		isInterleaving = true
	} else if _s2[j] == _s3[k] && dp(i, j-1, k-1) {
		isInterleaving = true
	}

	memo[state] = isInterleaving
	return
}

// @lc code=end

