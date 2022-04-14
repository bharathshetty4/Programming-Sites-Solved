package main

/*
URL: https://leetcode.com/problems/search-in-a-binary-search-tree
Status: Success
Runtime: 22 ms, faster than 74.25% of Go online submissions for Search in a Binary Search Tree.
Memory Usage: 7 MB, less than 95.81% of Go online submissions for Search in a Binary Search Tree.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func searchBST(root *TreeNode, val int) *TreeNode {
    
    if root == nil {
        return nil
    }
    if root.Val == val{
        return root
    }
    lTree := searchBST(root.Left, val)
    if lTree != nil {
        return lTree
    }
    rTree := searchBST(root.Right, val)
    if rTree != nil {
        return rTree
    }
    return nil
}
