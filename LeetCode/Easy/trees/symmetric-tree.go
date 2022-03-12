package main

/*
URL: https://leetcode.com/problems/symmetric-tree
Status: Success
Runtime: 6 ms, faster than 27.31% of Go online submissions for Symmetric Tree.
Memory Usage: 2.9 MB, less than 87.86% of Go online submissions for Symmetric Tree.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return symmetryCheck(root.Left, root.Right)
}

func symmetryCheck(node1, node2 *TreeNode) bool{
    if node1 == nil && node2 == nil{
        return true
    }
    if node1 == nil || node2 == nil {
        return false
    }
    if node1.Val != node2.Val{
        return false
    }
    
    
    return symmetryCheck(node1.Left, node2.Right)&&symmetryCheck(node1.Right, node2.Left)
}
