package main

/*
URL: https://leetcode.com/problems/linked-list-cycle
Runtime: 11 ms, faster than 60.08% of Go online submissions for Linked List Cycle.
Memory Usage: 4.3 MB, less than 88.57% of Go online submissions for Linked List Cycle.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	ptr1 := head
	ptr2 := head
	for {
		if ptr1 == nil || ptr1.Next == nil || ptr2.Next == nil || ptr2.Next.Next == nil {
			return false
		}
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next.Next
		// need to check the node itself as the duplicate values are allowed
		if ptr1 == ptr2 {
			return true
		}
	}
}
