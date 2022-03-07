package main

import (
   "fmt"
)

func main() {
	a := 1
	incrInt(&a)
	incrInt(&a)
	fmt.Println("a= ", a)
}

func incrInt(d *int) {
	*d++
}

// ptr by reference for int array
func inorderTraversal(root *TreeNode) []int {
    var list []int
    inorder(root,&list)
    return list
}

func inorder(node *TreeNode, list * []int) {
    if node == nil {
        return
    }
    inorder(node.Left, list)
    *list = append(*list,node.Val)
    inorder(node.Right, list)
}

// OR this works too
func inorderTraversal(root *TreeNode) []int {
    var list = new([]int)
    inorder(root,list)
    return *list
}

func inorder(node *TreeNode, list * []int) {
    if node == nil {
        return
    }
    inorder(node.Left, list)
    *list = append(*list,node.Val)
    inorder(node.Right, list)
}
