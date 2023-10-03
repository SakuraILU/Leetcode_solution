/*
 * @lc app=leetcode id=1048 lang=golang
 *
 * [1048] Longest String Chain
 */

// @lc code=start
var memo map[int]int

var _words []string

func longestStrChain(words []string) int {
	_words = words

	memo = make(map[int]int, 0)

	mlen := 0
	for i := 0; i < len(_words); i++ {
		mlen = max(mlen, dp(i))
	}

	return mlen
}

func dp(pos int) (mlen int) {
	if v, ok := memo[pos]; ok {
		return v
	}

	for i := 0; i < len(_words); i++ {
		if isPredecessor(_words[i], _words[pos]) {
			mlen = max(mlen, dp(i))
		}
	}
	mlen += 1

	memo[pos] = mlen
	return mlen
}

func isPredecessor(worda, wordb string) bool {
	if len(worda) != len(wordb)-1 {
		return false
	}
	first, second := 0, 0

	insert_cnt := 0
	for first < len(worda) && second < len(wordb) {
		if worda[first] != wordb[second] {
			second++
			insert_cnt++
			if insert_cnt > 1 {
				return false
			}
		} else {
			first++
			second++
		}
	}

	return insert_cnt == 1 || first == second
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
