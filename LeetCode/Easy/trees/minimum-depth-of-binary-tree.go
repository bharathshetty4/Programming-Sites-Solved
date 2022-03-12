package main

/*
URL: https://leetcode.com/problems/minimum-depth-of-binary-tree
Status: Success
Runtime: 287 ms, faster than 50.76% of Go online submissions for Minimum Depth of Binary Tree.
Memory Usage: 26.3 MB, less than 12.64% of Go online submissions for Minimum Depth of Binary Tree.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var (
        treeLevel int
        q []*TreeNode
    )
    q = append(q,root)
    for len(q) >0 {
        parentNodes := len(q) // total number of parent nodes of current level
        for parentNodes > 0 && len(q) >0{
            currNode := q[0]
            q = q[1:] // remove the current node from the queue
            if currNode.Left == nil && currNode.Right == nil {
                return treeLevel+1
            }
            if currNode.Left != nil{
                 q = append(q,currNode.Left)
            }
            if currNode.Right != nil{
                 q = append(q,currNode.Right)

            }
            parentNodes-- // remove the parent node count
        }
            treeLevel++ // increase the level by 1
    }
    return treeLevel
}
