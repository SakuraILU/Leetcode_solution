/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
// package leetcodesolution

// type Node struct {
// 	Val    int
// 	Next   *Node
// 	Random *Node
// }

var org2new map[*Node]*Node

func copyRandomList(head *Node) *Node {
	org2new = make(map[*Node]*Node, 0)

	return dfs(head)
}

func dfs(head *Node) (new_head *Node) {
	if head == nil {
		return
	}

	new_head = &Node{Val: head.Val}
	org2new[head] = new_head

	new_head.Next = dfs(head.Next)
	new_head.Random = org2new[head.Random]

	return new_head
}

// @lc code=end
