/*
 * @lc app=leetcode id=790 lang=golang
 *
 * [790] Domino and Tromino Tiling
 */

// @lc code=start
// package leetcodesolution

// o    x     o
// o -- x --> o

// o    x      o
// o -- xx --> x

// o    xx     x
// o -- x  --> o

// o    xx     o
// o -- xx --> o (next.next)

// o    xx     x
// x --    --> o

// o    xx     o
// x --  x --> o (next.next)

// x           o
// o -- xx --> x

// x	  x     o
// o -- xx --> o (next.next)

type State struct {
	ncol   int
	column Column
}

var memo map[State]int

type Column int

func (s Column) String() string {
	switch s {
	case OO:
		return "OO"
	case OX:
		return "OX"
	case XO:
		return "XO"
	}

	return "Unknown"
}

const (
	OO Column = iota
	XO
	OX
)

var _n int
var M int

func numTilings(n int) int {
	_n = n
	M = int(math.Pow(10, 9) + 7)
	memo = make(map[State]int)
	return dp(0, OO)
}

func dp(ncol int, column Column) (res int) {
	if ncol == _n {
		if column == OO {
			return 1
		} else {
			return 0
		}
	} else if ncol > _n {
		return 0
	}

	state := State{ncol: ncol, column: column}
	if v, ok := memo[state]; ok {
		return v
	}

	switch column {
	case OO:
		res = dp(ncol+1, OO) + dp(ncol+1, OX) + dp(ncol+1, XO) + dp(ncol+2, OO)
	case OX:
		res = dp(ncol+1, XO) + dp(ncol+2, OO)
	case XO:
		res = dp(ncol+1, OX) + dp(ncol+2, OO)
	}

	res %= M
	memo[state] = res
	return res
}

// @lc code=end
