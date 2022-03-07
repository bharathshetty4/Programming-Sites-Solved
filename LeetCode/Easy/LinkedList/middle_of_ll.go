package main
/*
URL: https://leetcode.com/problems/middle-of-the-linked-list/submissions/
Runtime: 0 ms, faster than 100.00% of Go online submissions for Middle of the Linked List.
Memory Usage: 2 MB, less than 44.00% of Go online submissions for Middle of the Linked List.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
    if head == nil {
		return nil
	}
	ptr := head
	ptr2 := head
	for {
		if ptr2.Next == nil {
			return ptr
		}
		if ptr2.Next.Next == nil {
			return ptr.Next
		}
		ptr2 = ptr2.Next.Next
		ptr = ptr.Next
	}
}
