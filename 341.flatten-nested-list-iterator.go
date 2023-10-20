/*
 * @lc app=leetcode id=341 lang=golang
 *
 * [341] Flatten Nested List Iterator
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

// package leetcodesolution

// type NestedInteger struct {
// }

// func (this NestedInteger) IsInteger() bool           { return true }
// func (this NestedInteger) GetInteger() int           { return 0 }
// func (n *NestedInteger) SetInteger(value int)        {}
// func (this NestedInteger) GetList() []*NestedInteger { return make([]*NestedInteger, 0) }

type NestedIterator struct {
	stack [][]*NestedInteger
	next  int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	ni := &NestedIterator{
		stack: make([][]*NestedInteger, 0),
	}

	ni.stack = append(ni.stack, nestedList)

	return ni
}

func (this *NestedIterator) Next() (val int) {
	return this.next
}

func (this *NestedIterator) HasNext() bool {
dfs:
	if len(this.stack) == 0 {
		return false
	}
	// pop a list
	lst := this.stack[len(this.stack)-1]
	this.stack = this.stack[0 : len(this.stack)-1]

	for len(lst) > 0 && !lst[0].IsInteger() {
		// if has something left
		if len(lst) > 1 {
			this.stack = append(this.stack, lst[1:len(lst)])
		}

		lst = lst[0].GetList()
	}

	if len(lst) == 0 {
		goto dfs
	}

	this.next = lst[0].GetInteger()
	if len(lst) > 1 {
		this.stack = append(this.stack, lst[1:len(lst)])
	}

	return true
}

// @lc code=end
