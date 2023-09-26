/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	new_head, successor := reverseN(head, k)
	// not reversed, nums are not enough... just return
	if new_head == nil {
		return head
	}

	head.Next = reverseKGroup(successor, k)

	return new_head
}

func reverseN(head *ListNode, n int) (new_head *ListNode, successor *ListNode) {
	// enough nums
	if n == 1 {
		return head, head.Next
	}
	// nums of nodes < need nums...
	// do nothing
	if head.Next == nil {
		return nil, nil
	}
	new_head, successor = reverseN(head.Next, n-1)
	if new_head != nil {
		head.Next.Next = head
	}

	return new_head, successor
}

// @lc code=end
