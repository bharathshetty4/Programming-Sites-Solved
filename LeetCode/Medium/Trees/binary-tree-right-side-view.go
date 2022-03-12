/*
URL: https://leetcode.com/problems/binary-tree-right-side-view/
Status: Success
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Right Side View.
Memory Usage: 2.3 MB, less than 79.45% of Go online submissions for Binary Tree Right Side View.
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

func rightSideView(root *TreeNode) []int {
    var (
        tree []int
		q []*TreeNode
	)
    if root == nil {
        return []int{}
    }
    currNode := root
	q = append(q, root)
	for len(q) > 0 {
		parentNodes := len(q) // total number of parent nodes of current level
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
		}
        tree = append(tree, currNode.Val)
	}
    return tree
}
