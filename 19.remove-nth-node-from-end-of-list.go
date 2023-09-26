/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var _n, k int
var prenode *ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	_n = n
	k = 0
	prenode = nil

	traverse(head)

	if prenode == nil {
		head = head.Next
	} else {
		prenode.Next = prenode.Next.Next
	}

	return head
}

func traverse(node *ListNode) {
	if node == nil {
		return
	}

	traverse(node.Next)
	k++

	if k == _n+1 {
		fmt.Println("find one")
		prenode = node
	}
}

// @lc code=end

