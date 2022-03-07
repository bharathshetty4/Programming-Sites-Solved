package main

/*
URL: https://leetcode.com/problems/binary-tree-inorder-traversal
Status: Success
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
Memory Usage: 1.9 MB, less than 91.64% of Go online submissions for Binary Tree Inorder Traversal.
*/

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var list = new([]int)
    inorder(root,list)
    return *list
}

func inorder(node *TreeNode, list * []int) {
    if node == nil {
        return
    }
    inorder(node.Left, list)
    *list = append(*list,node.Val)
    inorder(node.Right, list)
}

