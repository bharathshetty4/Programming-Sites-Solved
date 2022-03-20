/*
URL: https://leetcode.com/problems/binary-tree-level-order-traversal
Status: Success
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal.
Memory Usage: 2.8 MB, less than 51.13% of Go online submissions for Binary Tree Level Order Traversal.
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    var (
		q []*TreeNode
        tree = [][]int{}
        levelTree []int
	)
    if root == nil {
        return [][]int{}
    }
    currNode := root
	q = append(q, root)
	for len(q) > 0 {
		parentNodes := len(q) // total number of parent nodes of current level
        levelTree=nil
		for parentNodes > 0 && len(q) > 0 {
			currNode = q[0]
			q = q[1:] // remove the current node from the queue
			
			if currNode.Left != nil {
				q = append(q, currNode.Left)
			}
			if currNode.Right != nil {
				q = append(q, currNode.Right)
			}
			parentNodes-- // remove the parent node count
            levelTree=append(levelTree,currNode.Val)
		}
        tree = append(tree,levelTree)
	}
    return tree
}
