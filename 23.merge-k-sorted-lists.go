/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
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

// import (
// 	"container/heap"
// )

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

type NodeHeap struct {
	nodes []*ListNode
}

func NewNodeHeap() *NodeHeap {
	return &NodeHeap{
		nodes: make([]*ListNode, 0),
	}
}

func (n *NodeHeap) Push(v interface{}) {
	n.nodes = append(n.nodes, v.(*ListNode))
}

func (n *NodeHeap) Less(i, j int) bool {
	return n.nodes[i].Val < n.nodes[j].Val
}

func (n *NodeHeap) Swap(i, j int) {
	tmp := n.nodes[i]
	n.nodes[i] = n.nodes[j]
	n.nodes[j] = tmp
}

func (n *NodeHeap) Pop() interface{} {
	v := n.nodes[len(n.nodes)-1]
	n.nodes = n.nodes[0 : len(n.nodes)-1]
	return v
}

func (n *NodeHeap) Len() int {
	return len(n.nodes)
}

func mergeKLists(lists []*ListNode) *ListNode {
	nheap := NewNodeHeap()
	heap.Init(nheap)

	for _, head := range lists {
		if head != nil {
			heap.Push(nheap, head)
		}
	}

	var head, tail *ListNode
	for nheap.Len() > 0 {
		mnode := heap.Pop(nheap).(*ListNode)

		if head == nil {
			head = mnode
			tail = head
		} else {
			tail.Next = mnode
			tail = tail.Next
		}

		if nnode := mnode.Next; nnode != nil {
			heap.Push(nheap, nnode)
		}
	}

	return head
}

// @lc code=end
