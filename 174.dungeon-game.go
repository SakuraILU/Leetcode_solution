/*
 * @lc app=leetcode id=174 lang=golang
 *
 * [174] Dungeon Game
 */

// @lc code=start
type State struct {
	i, j int
}

var memo map[State]int

var _dungeon [][]int

func calculateMinimumHP(dungeon [][]int) int {
	_dungeon = dungeon

	memo = make(map[State]int)
	return dp(0, 0)
}

func dp(i, j int) (mhp int) {
	if i == len(_dungeon)-1 && j >= len(_dungeon[0]) {
		return 1
	} else if i >= len(_dungeon) && j == len(_dungeon[0])-1 {
		return 1
	} else if i >= len(_dungeon) || j >= len(_dungeon[0]) {
		return math.MaxInt
	}

	state := State{i: i, j: j}
	if v, ok := memo[state]; ok {
		return v
	}

	mhp = -_dungeon[i][j] + min(dp(i+1, j), dp(i, j+1))

	// at least has one hp...
	// i.e. (i+1, j) need 5 hp at least, and there are 10 hp supply packs..
	// 		(i, j) only need -5 hp XD, but of course the hero is dead if hp == -5,
	//		so, at least hero need 1 hp to gurantee its alive
	if mhp < 1 {
		mhp = 1
	}

	memo[state] = mhp
	return
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end
