/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var _head *ListNode
var _k int
var rank, num int
var new_head *ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	_head = head
	_k = k
	rank, num = 0, 0
	new_head = head

	traverse(head)

	return new_head
}

func traverse(node *ListNode) {
	if node == nil {
		return
	}

	num++
	traverse(node.Next)
	rank++
	// rotate k rounds, nothing need to do...
	if _k%num == 0 {
		return
	}

	// last node -> first node
	if rank == 1 {
		node.Next = _head
	}
	if rank == _k%num {
		// record new head
		new_head = node
	} else if rank == _k%num+1 {
		// adjust the new last tail -> nil
		node.Next = nil
	}
}

// @lc code=end
