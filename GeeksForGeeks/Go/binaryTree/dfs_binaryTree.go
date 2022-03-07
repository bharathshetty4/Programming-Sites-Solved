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

// using stack
func preorderTraversal(root *TreeNode) []int {
    var ans []int
    if root == nil {
        return ans
    }
    stack := []*TreeNode{root}
    for len(stack) > 0 {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        ans = append(ans, top.Val)
        if top.Right != nil {
            stack = append(stack, top.Right)
        }
        if top.Left != nil {
            stack = append(stack, top.Left)
        }
    }
    return ans
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
