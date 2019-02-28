//https://www.hackerrank.com/challenges/reverse-a-doubly-linked-list/problem?h_l=interview&playlist_slugs%5B%5D%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D%5B%5D=linked-lists
package main

import (
	"fmt"
)

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

func (doublyLinkedList *DoublyLinkedList) insertNodeIntoDoublyLinkedList(nodeData int32) {
	node := &DoublyLinkedListNode{
		next: nil,
		prev: nil,
		data: nodeData,
	}

	if doublyLinkedList.head == nil {
		doublyLinkedList.head = node
	} else {
		doublyLinkedList.tail.next = node
		node.prev = doublyLinkedList.tail
	}

	doublyLinkedList.tail = node
}

func printLinkedList(head *DoublyLinkedListNode) {
	node := head
	for {
		if node == nil {
			fmt.Println()
			return
		}
		fmt.Printf("%d->", node.data)
		node = node.next
	}
}

// Complete the reverse function below.

/*
 * For your reference:
 *
 * DoublyLinkedListNode {
 *     data int32
 *     next *DoublyLinkedListNode
 *     prev *DoublyLinkedListNode
 * }
 *
 */
func reverse(head *DoublyLinkedListNode) *DoublyLinkedListNode {
	var next *DoublyLinkedListNode
	cur := head
	for {
		if cur == nil {
			return cur
		}
		next = cur.next

		cur.next = cur.prev
		cur.prev = next

		if next == nil {
			return cur
		}
		cur = next
	}

	return cur
}

func main() {
	llist := DoublyLinkedList{}
	llist.insertNodeIntoDoublyLinkedList(1)
	llist.insertNodeIntoDoublyLinkedList(4)
	llist.insertNodeIntoDoublyLinkedList(5)
	printLinkedList(llist.head)
	llist1 := reverse(llist.head)
	printLinkedList(llist1)
}
