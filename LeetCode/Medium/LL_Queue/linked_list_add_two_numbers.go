package main

/*
URL: https://leetcode.com/problems/add-two-numbers/
Status: Passed
Runtime: 8 ms, faster than 93.28% of Go online submissions for Add Two Numbers.
Memory Usage: 5.1 MB, less than 22.07% of Go online submissions for Add Two Numbers.
*/
import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	listSum := make([]int, 0)
	reminder := 0
	tmp := l1
	for {
		if l1 == nil && l2 == nil && reminder == 0 {
			break
		}
		sum := reminder
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}

		if sum >= 10 {
			listSum = append(listSum, sum%10)
			reminder = int(sum / 10)
		} else {
			reminder = 0
			listSum = append(listSum, sum)
		}

		if l1 != nil && l1.Next != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}
		if l2 != nil && l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
	}
	l1 = tmp
	lenSum := len(listSum)

	for key, val := range listSum {
		l1.Val = val
		if key == lenSum-1 {
			break
		}
		if l1.Next == nil {
			l1.Next = &ListNode{}
		}
		l1 = l1.Next
	}

	return tmp
}

func main() {
	l13 := &ListNode{3, nil}
	l12 := &ListNode{4, l13}
	l11 := &ListNode{2, l12}

	l23 := &ListNode{4, nil}
	l22 := &ListNode{6, l23}
	l21 := &ListNode{5, l22}

	list := addTwoNumbers(l11, l21)

	fmt.Println("list", list)
}
