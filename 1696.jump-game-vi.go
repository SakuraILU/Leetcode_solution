/*
 * @lc app=leetcode id=1696 lang=golang
 *
 * [1696] Jump Game VI
 */

// @lc code=start
package main

import (
	"container/list"
	"math"
)

type Entry struct {
	idx int
	val int
}

type Queue struct {
	size int
	lst  list.List
}

func NewQueue(size int) *Queue {
	q := &Queue{
		size: size,
		lst:  *list.New(),
	}
	// i don't know why...but must do this
	// otherwise, after several PushFront(), you will see
	// there is a elem with <nil> value in the list
	// when you do traverse or Back()...really weired
	q.lst.Init()
	return q
}

func (q *Queue) Push(idx, val int) {
	if elem := q.lst.Back(); elem != nil {
		entry := elem.Value.(*Entry)
		if entry.idx >= idx+q.size {
			q.lst.Remove(elem)
		}
	}

	for {
		elem := q.lst.Front()
		if elem == nil {
			break
		}

		entry := elem.Value.(*Entry)
		if entry.val < val {
			q.lst.Remove(elem)
		} else {
			break
		}
	}

	q.lst.PushFront(&Entry{idx: idx, val: val})
}

func (q *Queue) Max() int {
	return q.lst.Back().Value.(*Entry).val
}

const INVALID = math.MaxInt

func maxResult(nums []int, k int) int {
	length := len(nums)

	memo := make([]int, length)
	memo[length-1] = nums[length-1]
	queue := NewQueue(k)
	queue.Push(length-1, nums[length-1])
	for i := length - 2; i >= 0; i-- {
		memo[i] = nums[i] + queue.Max()
		queue.Push(i, memo[i])
	}

	return memo[0]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
