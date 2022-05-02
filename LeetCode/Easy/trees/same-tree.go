package main

/*
URL: https://leetcode.com/problems/same-tree
Status: Success
Runtime: 4 ms, faster than 12.40% of Go online submissions for Same Tree.
Memory Usage: 2.1 MB, less than 69.36% of Go online submissions for Same Tree.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p==nil || q==nil || p.Val!=q.Val{
        return false
    }
    
    return isSameTree(p.Left,q.Left)&&isSameTree(p.Right,q.Right)
    
}
