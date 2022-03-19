package main

/*
URL: https://leetcode.com/problems/palindrome-linked-list
Runtime: 224 ms, faster than 49.76% of Go online submissions for Palindrome Linked List.
Memory Usage: 9.4 MB, less than 90.62% of Go online submissions for Palindrome Linked List.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }
    middle := getMiddle(head)
    reversedMiddle := reverseLinkedList(middle)
    for {
        if reversedMiddle == nil || head == nil {
            return true
        }
        if reversedMiddle.Val != head.Val {
            return false
        }
        reversedMiddle = reversedMiddle.Next
        head =head.Next
    }
}

func getMiddle(head *ListNode) *ListNode{
    ptr := head
	ptr2 := head
	for {
		if ptr2.Next == nil {
			return  ptr
		}
		if ptr2.Next.Next == nil {
			return ptr
		}
		ptr2 = ptr2.Next.Next
		ptr = ptr.Next
	}
}

func reverseLinkedList(head *ListNode) *ListNode {
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
