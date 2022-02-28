package main

/*
URL: https://leetcode.com/problems/swap-nodes-in-pairs
Status: Success
Runtime: 0 ms, faster than 100.00% of Go online submissions for Swap Nodes in Pairs.
Memory Usage: 2.1 MB, less than 20.58% of Go online submissions for Swap Nodes in Pairs.
*/
import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    newHead := head
    curr := head
    var prev, next *ListNode 
    if head.Next != nil {
        newHead=head.Next
    }
    for {
        if curr == nil || curr.Next == nil {
            break
        }
        next = curr.Next
        if prev != nil {
            prev.Next = next
        }
        curr.Next = next.Next
        next.Next = curr
        prev = curr
        curr = curr.Next
        
    }
    
    return newHead
}
