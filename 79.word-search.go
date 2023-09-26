/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
// package leetcodesolution

type State struct {
	i, j int
}

var choices = []State{State{-1, 0}, State{1, 0}, State{0, -1}, State{0, 1}}

var gboard [][]byte
var gword string
var visited [][]bool

func exist(board [][]byte, word string) bool {
	gboard = board
	gword = word
	visited = make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(i, j, word_i int) bool {
	if word_i >= len(gword) {
		return true
	} else if i < 0 || i >= len(gboard) || j < 0 || j >= len(gboard[0]) {
		return false
	}

	if visited[i][j] {
		return false
	}

	if gboard[i][j] != gword[word_i] {
		return false
	}

	visited[i][j] = true
	for _, choice := range choices {
		ni, nj := i+choice.i, j+choice.j

		if dfs(ni, nj, word_i+1) {
			return true
		}
	}
	visited[i][j] = false

	return false
}

// @lc code=end
