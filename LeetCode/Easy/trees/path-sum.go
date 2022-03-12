package main

/*
URL: https://leetcode.com/problems/path-sum
Status: Success
Runtime: 4 ms, faster than 87.02% of Go online submissions for Path Sum.
Memory Usage: 4.6 MB, less than 94.84% of Go online submissions for Path Sum.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    if root.Val == targetSum && root.Left == nil && root.Right == nil {
        return true
    }
    targetSum = targetSum - root.Val
    return hasPathSum(root.Left,targetSum)||hasPathSum(root.Right,targetSum)
}
