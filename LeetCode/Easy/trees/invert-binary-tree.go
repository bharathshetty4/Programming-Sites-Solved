package main

/*
URL: https://leetcode.com/problems/invert-binary-tree/submissions/
Status: Success
Runtime: 0 ms, faster than 100.00% of Go online submissions for Invert Binary Tree.
Memory Usage: 2.1 MB, less than 72.10% of Go online submissions for Invert Binary Tree.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

func invert(node *TreeNode) {
	if node == nil {
		return
	}
	r := node.Right
	node.Right = node.Left
	node.Left = r
	invert(node.Left)
	invert(node.Right)
}
