/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var first *ListNode
var newhead, newtail *ListNode
var over bool

func reorderList(head *ListNode) {
	first = head
	over = false

	dfs(head)

	return
}

func dfs(node *ListNode) {
	if node == nil {
		return
	}

	dfs(node.Next)

	if over {
		return
	}

	// check the last pair to concat
	if first == node || first.Next == node {
		over = true
	}

	// a thing to notice:
	// if first == node,
	// we can still concat it (first) and concat it (node) again, doesn't matter
	// for the second concatenation, newtail.Next --> itself and traverse to itself again...

	// concat first
	if newhead == nil {
		newhead = first
		newtail = first
	} else {
		newtail.Next = first
		newtail = newtail.Next
	}
	first = first.Next

	// concat second
	newtail.Next = node
	newtail = newtail.Next
	newtail.Next = nil
}

// @lc code=end
