package main

/*
URL: https://leetcode.com/problems/populating-next-right-pointers-in-each-node/submissions/
Status: Success
Runtime: 16 ms, faster than 10.63% of Go online submissions for Populating Next Right Pointers in Each Node.
Memory Usage: 6.8 MB, less than 49.21% of Go online submissions for Populating Next Right Pointers in Each Node.
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

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var (
		q []*Node
	)
	q = append(q, root)
	for len(q) > 0 {
		parentNodes := len(q) // total number of parent nodes of current level
		var prevNode *Node
		prevNode = nil
		for parentNodes > 0 && len(q) > 0 {
			currNode := q[0]
			q = q[1:] // remove the current node from the queue
			if prevNode != nil {
				prevNode.Next = currNode
			}
			if currNode.Left != nil {
				q = append(q, currNode.Left)
			}
			if currNode.Right != nil {
				q = append(q, currNode.Right)
			}
			prevNode = currNode
			parentNodes-- // remove the parent node count
		}
	}
	return root
}

// Best Answer
func connectBest(root *Node) *Node {
    if root == nil {
        return root
    }
    connectNode(root.Left, root.Right)
    return root
}
func connectNode(left *Node, right *Node) {
    if left == nil {
        return
    }
    left.Next = right
    connectNode(left.Left, left.Right)
    connectNode(right.Left, right.Right)
    connectNode(left.Right, right.Left)
}
