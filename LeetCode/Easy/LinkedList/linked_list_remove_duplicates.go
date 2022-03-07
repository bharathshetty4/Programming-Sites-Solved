package main
/*
URL: https://leetcode.com/problems/remove-duplicates-from-sorted-list
Runtime: 9 ms, faster than 19.51% of Go online submissions for Remove Duplicates from Sorted List.
Memory Usage: 3.1 MB, less than 75.73% of Go online submissions for Remove Duplicates from Sorted List.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    currentNode := head
	for {
		if currentNode == nil || currentNode.Next == nil {
			return head
		}
		if currentNode.Val == currentNode.Next.Val {
			currentNode.Next = currentNode.Next.Next
			continue
		}
		currentNode = currentNode.Next
	}
}
