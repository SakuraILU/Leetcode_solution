/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head

	slowpre := slow
	for fast != nil && fast.Next != nil {
		slowpre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	slowpre.Next = nil

	head = sortList(head)
	slow = sortList(slow)

	return merge(head, slow)
}

func merge(head1, head2 *ListNode) (head *ListNode) {
	tail := head
	for head1 != nil && head2 != nil {
		var tmp *ListNode
		if head1.Val <= head2.Val {
			tmp = head1
			head1 = head1.Next
		} else {
			tmp = head2
			head2 = head2.Next
		}

		if head == nil {
			head = tmp
			tail = tmp
		} else {
			tail.Next = tmp
			tail = tail.Next
		}
		tail.Next = nil
	}

	if head1 != nil {
		tail.Next = head1
	} else {
		tail.Next = head2
	}

	return
}

// @lc code=end
