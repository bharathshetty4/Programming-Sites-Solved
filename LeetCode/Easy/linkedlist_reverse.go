package main
/*
URL: https://leetcode.com/problems/reverse-linked-list
Runtime: 6 ms, faster than 20.48% of Go online submissions for Reverse Linked List.
Memory Usage: 2.6 MB, less than 93.28% of Go online submissions for Reverse Linked List.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	currentNode := head
	for {
		if currentNode == nil {
			return prev
		}
		next = currentNode.Next
		currentNode.Next = prev
		prev = currentNode
		currentNode = next
	}
}
