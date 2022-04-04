package main

/*
URL: https://leetcode.com/problems/swapping-nodes-in-a-linked-list/submissions/
Status: Accepted
Runtime: 183 ms, faster than 83.33% of Go online submissions for Swapping Nodes in a Linked List.
Memory Usage: 10.4 MB, less than 34.37% of Go online submissions for Swapping Nodes in a Linked List.
*/



func swapNodes(head *ListNode, k int) *ListNode {
    if head == nil {
		return nil
	}
    // handle if it is head
    var firstNode, lastNode *ListNode
	ptr := head
	ptr2 := head
    nodeLen :=0
	for {
        nodeLen++
        if nodeLen == k{
            firstNode = ptr
        }
		if ptr2.Next == nil {
            nodeLen = (2*nodeLen)-1
			break
		}
		if ptr2.Next.Next == nil {
            nodeLen= (2*nodeLen) // extra len as it has one more node
			break
		}
		ptr2 = ptr2.Next.Next
		ptr = ptr.Next
	}
    lastNodeIndex := nodeLen - k + 1
    ptr = head
    idxCnt := 0 
    for firstNode== nil || lastNode == nil{
        if ptr == nil {
            // some error, return as it is
            return head
        }
        idxCnt++
        if idxCnt == k {
            firstNode = ptr
        }
        if idxCnt == lastNodeIndex {
            lastNode = ptr
        }
        ptr = ptr.Next
    }
    firstNode.Val,lastNode.Val = lastNode.Val,firstNode.Val
    return head
}
