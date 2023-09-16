/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "fmt"

var targetsum int = 0
var cursum int = 0
var curpath []int
var res [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	targetsum = targetSum
	cursum = 0
	curpath = make([]int, 0)
	res = make([][]int, 0)

	traverse(root)
	return res
}

func traverse(node *TreeNode) {
	if node == nil {
		fmt.Println("return from nul")
		return
	}

	curpath = append(curpath, node.Val)
	cursum += node.Val

	// fmt.Printf("Goto node %v: %v list[%v]\n", node.Val, cursum, curpath)

	if node.Left == nil && node.Right == nil {
		if cursum == targetsum {
			// fmt.Printf("get a result %v\n", curpath)
			path := make([]int, len(curpath))
			copy(path, curpath)
			res = append(res, path)
		}
	}

	traverse(node.Left)
	traverse(node.Right)
	cursum -= node.Val
	curpath = curpath[:len(curpath)-1]
}

// @lc code=end
