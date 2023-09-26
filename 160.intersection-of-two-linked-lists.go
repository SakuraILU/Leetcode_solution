/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// package leetcodesolution

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0

	itrA, itrB := headA, headB
	for itrA != nil {
		lenA++
		itrA = itrA.Next
	}
	for itrB != nil {
		lenB++
		itrB = itrB.Next
	}

	itrA, itrB = headA, headB
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			itrA = itrA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			itrB = itrB.Next
		}
	}

	for itrA != itrB {
		itrA = itrA.Next
		itrB = itrB.Next
	}

	return itrA
}

// @lc code=end
