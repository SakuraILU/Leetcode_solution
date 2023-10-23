/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
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

// package leetcodesolution

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

const (
	NULL = "#"
	SEP  = " "
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	stream := this.ser(root)
	return stream
}

func (this *Codec) ser(root *TreeNode) string {
	if root == nil {
		return NULL
	}

	ldata := this.ser(root.Left)
	rdata := this.ser(root.Right)

	return strconv.Itoa(root.Val) + SEP + ldata + SEP + rdata
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(stream string) *TreeNode {
	data := strings.Split(stream, SEP)
	return this.dec(&data)
}

func (this *Codec) dec(pdata *[]string) *TreeNode {
	val := (*pdata)[0]
	(*pdata) = (*pdata)[1:len((*pdata))]

	if val == NULL {
		return nil
	}

	rootval, _ := strconv.Atoi(val)
	root := &TreeNode{
		Val:   rootval,
		Left:  this.dec(pdata),
		Right: this.dec(pdata),
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end
