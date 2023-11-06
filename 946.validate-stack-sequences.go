/*
 * @lc app=leetcode id=946 lang=golang
 *
 * [946] Validate Stack Sequences
 */

// @lc code=start
var stack []int
var _pushed, _popped []int

func validateStackSequences(pushed []int, popped []int) bool {
	stack = make([]int, 0)
	_pushed, _popped = pushed, popped

	return dfs(0, 0)
}

func dfs(i, j int) bool {
	if j == len(_popped) {
		return true
	}

	if len(stack) > 0 && stack[len(stack)-1] == _popped[j] {
		stack = stack[0 : len(stack)-1]
		return dfs(i, j+1)
	} else if i < len(_pushed) {
		stack = append(stack, _pushed[i])
		return dfs(i+1, j)
	} else {
		return false
	}
}

// @lc code=end
