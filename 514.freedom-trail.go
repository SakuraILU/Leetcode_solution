/*
 * @lc app=leetcode id=514 lang=golang
 *
 * [514] Freedom Trail
 */

// @lc code=start

type State struct {
	i, j int
}

var memo map[State]int

var _ring, _key string

func findRotateSteps(ring string, key string) int {
	_ring, _key = ring, key

	memo = make(map[State]int)
	return dp(0, 0)
}

func dp(i, j int) (mcnt int) {
	if j >= len(_key) {
		return 0
	}

	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	mcnt = math.MaxInt
	if _ring[i] == _key[j] {
		mcnt = 1 + dp(i, j+1)
	} else {
		for k := 0; k < len(_ring); k++ {
			if _ring[k] == _key[j] {
				delta := abs(k - i)
				mrotcnt := min(delta, len(_ring)-delta)
				mcnt = min(mcnt, 1+mrotcnt+dp(k, j+1))
			}
		}
	}

	memo[state] = mcnt
	return
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

// @lc code=end
