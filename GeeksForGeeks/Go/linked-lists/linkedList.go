package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func PrintLinkedList(head *Node) {
	if head == nil {
		fmt.Println("Linked List Head is nil")
	}
	node := head
	for {
		fmt.Print(node.data, "->")
		if node.next == nil {
			break
		}
		node = node.next
	}
	fmt.Println("nil")
}

func AddNode(head *Node, value int) *Node {
	newNode := new(Node)
	newNode.data = value
	if head == nil {
		head = newNode
		return head
	}

	node := head
	for {
		if node.next == nil {
			break
		}
		node = node.next
	}
	node.next = newNode
	return head
}

func DeleteNode(head *Node, position int) *Node {
	if head == nil {
		fmt.Println("The Linked list head is nil")
		return head
	}
	if position == 0 {
		head = nil
		return head
	}

	nodeRef := head
	i := 0
	for {
		lastNodeRef := nodeRef
		if nodeRef.next != nil {
			nodeRef = nodeRef.next
			i++
		} else {
			fmt.Println("Node Not found for position: ", position)
			return head
		}

		if i == position {
			fmt.Println("Deleted the Node at position: ", position)
			lastNodeRef.next = nodeRef.next
			return head
		}
	}

	return head
}

func main() {
	var head *Node
	head = AddNode(head, 20)
	AddNode(head, 30)
	AddNode(head, 40)
	AddNode(head, 50)
	// PrintLinkedList(head)
	// RemoveDuplicates(head)
	// head = DeleteNode(head, 1)
	// head = DeleteNode(head, 5)
	// PrintMiddleNode(head)
	// head = ReverseLinkedList(head)
	// head = ReverseWithGroup(head, 2)
	/* var head1 *Node;head1 = AddNode(head1, 22);AddNode(head1, 23);AddNode(head1, 24);ListAlternate(head, head1) */
	PrintLinkedList(head)
}

// solutions to the g4g questions

// https://www.geeksforgeeks.org/write-a-c-function-to-print-the-middle-of-the-linked-list/
func PrintMiddleNode(head *Node) {
	if head == nil {
		fmt.Println("No LL")
		return
	}
	ptr := head
	ptr2 := head
	for {
		if ptr2.next == nil {
			fmt.Println("Middle Node", ptr.data)
			break
		}
		if ptr2.next.next == nil {
			fmt.Println("Middle Node", ptr.next.data)
			break
		}
		ptr2 = ptr2.next.next
		ptr = ptr.next
	}
}

// https://www.geeksforgeeks.org/reverse-a-linked-list/
func ReverseLinkedList(head *Node) *Node {
	var prev, next *Node
	currentNode := head
	for {
		if currentNode == nil {
			return prev
		}
		next = currentNode.next
		currentNode.next = prev
		prev = currentNode
		currentNode = next
	}
}

// https://practice.geeksforgeeks.org/problems/remove-duplicate-element-from-sorted-linked-list/1
func RemoveDuplicates(head *Node) {
	currentNode := head
	for {
		if currentNode == nil || currentNode.next == nil {
			return
		}
		if currentNode.data == currentNode.next.data {
			currentNode.next = currentNode.next.next
			continue
		}
		currentNode = currentNode.next
	}
}

// https://practice.geeksforgeeks.org/problems/reverse-a-linked-list-in-groups-of-given-size/1
func ReverseWithGroup(head *Node, group int) *Node {
	var prev, next *Node
	currentNode := head
	for {
		// same old reversal logic with group limit
		for i := 0; i < group; i++ {
			if currentNode == nil {
				break
			}
			next = currentNode.next
			currentNode.next = prev
			prev = currentNode
			currentNode = next
		}
		// If the next node exist, set the head.next (first node of current recursive list) with the next reversed group
		if next != nil {
			head.next = ReverseWithGroup(next, group)
		}
		// return the last node of the group, which is the first node after the reversing the current group
		return prev
	}
}

// https://www.geeksforgeeks.org/merge-a-linked-list-into-another-linked-list-at-alternate-positions/
func ListAlternate(head, head2 *Node) {
	cur1 := head
	cur2 := head2
	for {
		if (cur1 == nil && cur2 == nil) || cur2 == nil {
			return
		}
		if cur1 != nil && cur1.next == nil {
			cur1.next = cur2
			cur2 = cur2.next
			cur1 = cur1.next
		}
		next2 := cur2.next
		next1 := cur1.next
		cur1.next = cur2
		cur1.next.next = next1
		cur2 = next2
		cur1 = cur1.next.next
	}
}
