/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
//  */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return add(l1, l2, 0)
}

func add(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	val := carry
	if l1 != nil {
		val += l1.Val
	}
	if l2 != nil {
		val += l2.Val
	}

	node := &ListNode{}
	if val < 10 {
		node.Val = val
		carry = 0
	} else {
		node.Val = val - 10
		carry = 1
	}

	if l1 == nil && l2 == nil {
		node.Next = add(nil, nil, carry)
	} else if l1 == nil {
		node.Next = add(nil, l2.Next, carry)
	} else if l2 == nil {
		node.Next = add(l1.Next, nil, carry)
	} else {
		node.Next = add(l1.Next, l2.Next, carry)
	}

	return node
}

// @lc code=end
