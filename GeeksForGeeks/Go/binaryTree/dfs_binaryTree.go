package main

import "fmt"

type Node struct {
	data       int
	leftChild  *Node
	rightChild *Node
}

func newNode(data int) *Node {
	newNode := new(Node)
	newNode.data = data
	return newNode
}

// DFS
func CreateTree() *Node {
	//initialize the root
	root := newNode(1)

	root.leftChild = newNode(2)
	root.rightChild = newNode(3)

	root.leftChild.leftChild = newNode(4)
	root.leftChild.rightChild = newNode(5)

	return root
}

func InOrder(node *Node) {
	if node == nil {
		return
	}
	InOrder(node.leftChild)
	fmt.Printf("%d->", node.data)
	InOrder(node.rightChild)
}

func PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d->", node.data)
	PreOrder(node.leftChild)
	PreOrder(node.rightChild)
}

func PostOrder(node *Node) {
	if node == nil {
		return
	}
	PostOrder(node.leftChild)
	PostOrder(node.rightChild)
	fmt.Printf("%d->", node.data)
}

func main() {
	root := CreateTree()

	//print preorder
	fmt.Println("\nPre Order:")
	PreOrder(root)

	//print inorder
	fmt.Println("\n\nIn Order:")
	InOrder(root)

	//print postorder
	fmt.Println("\n\nPost Order:")
	PostOrder(root)
}
