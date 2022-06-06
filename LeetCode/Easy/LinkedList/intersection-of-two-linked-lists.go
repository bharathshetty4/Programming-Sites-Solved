package main

/*
URL: https://leetcode.com/problems/intersection-of-two-linked-lists/
Runtime: 31 ms, faster than 91.18% of Go online submissions for Intersection of Two Linked Lists.
Memory Usage: 6.8 MB, less than 91.18% of Go online submissions for Intersection of Two Linked Lists.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    lenA :=  getLen(headA)
    lenB := getLen(headB)
    for {
        if lenA ==lenB {
            break
        }
        if lenA > lenB{
            headA = headA.Next
            lenA--
        }
        if lenB > lenA {
            headB = headB.Next
            lenB--
        }
    }
    for {
        if headB == nil || headA == nil {
            return headA //should not happen, safe case
        }
        if headA == headB {
            return headA
        }
        headA = headA.Next
        headB = headB.Next
    }
    return headA
    
}

func getLen(head *ListNode) int {
    cnt := 0
    for {
    if head == nil {
        return cnt
    }
        head = head.Next
        cnt++
    }
    return cnt
}
