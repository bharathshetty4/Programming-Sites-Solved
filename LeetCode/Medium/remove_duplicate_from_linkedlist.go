package main

/*
URL:
Status:
Runtime: <> ms, faster than <> of Go online submissions for Add Two Numbers.
Memory Usage: <> MB, less than <> of Go online submissions for Add Two Numbers.
*/
import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	var prevNode *ListNode
	lastNonDupNodeVal := head.Val
	head = nil
	for {
		if node == nil || node.Next == nil {
			// this is needed to handle if the last node itself is duplicated
			if node != nil && node.Val == lastNonDupNodeVal {
				if prevNode != nil {
					prevNode.Next = nil
				}
			}
			if head == nil && node != nil && node.Val != lastNonDupNodeVal {
				return node
			}
			return head
		}

		// both node has the same value
		if node.Next.Val == node.Val {
			node = node.Next
			continue
		}

	}
}

func main() {
	// -1 > 1 > 3 > 3
	n4 := ListNode{Val: 3}             // 4th node
	n3 := ListNode{Val: 3, Next: &n4}  // 3rd node
	n2 := ListNode{Val: 1, Next: &n3}  // 2nd node
	n1 := ListNode{Val: -1, Next: &n2} // 1st node

	resNode := deleteDuplicates(&n1)
	for {
		if resNode == nil {
			fmt.Println("Nil node")
			return
		}
		fmt.Println("Node value", resNode.Val)
		if resNode.Next == nil {
			return
		}
		resNode = resNode.Next
	}
}
