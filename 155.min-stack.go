/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start

type MinStack struct {
	stack     []int
	min_stack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:     make([]int, 0),
		min_stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.min_stack) > 0 {
		premin := this.min_stack[len(this.min_stack)-1]
		this.min_stack = append(this.min_stack, min(premin, val))
	} else {
		this.min_stack = append(this.min_stack, val)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[0 : len(this.stack)-1]
	this.min_stack = this.min_stack[0 : len(this.min_stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min_stack[len(this.min_stack)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
