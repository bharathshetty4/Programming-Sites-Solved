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
	PrintLinkedList(head)

	head = AddNode(head, 30)
	PrintLinkedList(head)

	head = AddNode(head, 40)
	PrintLinkedList(head)

	head = DeleteNode(head, 1)
	PrintLinkedList(head)

	head = DeleteNode(head, 5)
	PrintLinkedList(head)
}
