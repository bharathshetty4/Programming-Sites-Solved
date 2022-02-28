package main

import "fmt"

// tree node structure
type Node struct {
	data       int
	leftChild  *Node
	rightChild *Node
}

type queue struct {
	queue []*Node
}

// add the element to the end
func (q *queue) enque(node *Node) {
	q.queue = append(q.queue, node)
}

// remove the first element and return it
func (q *queue) deque() *Node {
	if len(q.queue) <= 0 {
		return nil
	}
	node := q.queue[0]
	q.queue = q.queue[1:]
	return node
}

func newNode(data int) *Node {
	newNode := new(Node)
	newNode.data = data
	return newNode
}

// find the parent value and create a new node. New node is placed on parent node based on the leftChild value
// TODO: we can change the logic to insert to the next previous node too.
/*
 if (temp->left != NULL){
    q.push(temp->left);
else {
    temp->left = CreateNode(data);
    }}
*/
func addNode(root *Node, parentVal, newNodeVal int, leftChild bool) {
	node := newNode(newNodeVal)
	q := &queue{}
	q.enque(root)
	for {
		// keep on going until the queue is empty
		if len(q.queue) <= 0 {
			break
		}
		queueNode := q.deque() // pop
		// if the parent node is the node we are searching for,stop the op
		if queueNode.data == parentVal {
			if leftChild {
				queueNode.leftChild = &Node{data: newNodeVal}
				return
			}
			queueNode.rightChild = node
			return
		}

		if queueNode.leftChild != nil {
			q.enque(queueNode.leftChild)
		}
		if queueNode.rightChild != nil {
			q.enque(queueNode.rightChild)
		}
	}
}

func inOrder(node *Node) {
	if node == nil {
		return
	}
	inOrder(node.leftChild)
	fmt.Printf("%d->", node.data)
	inOrder(node.rightChild)
}

// keep on checking children sum for left and right child, until we hit the leaf.
// If one of the sum found to be not equal, return false
// https://www.geeksforgeeks.org/check-for-children-sum-property-in-a-binary-tree/
func checkForChildreSum(node *Node) bool {
	if node == nil || (node.leftChild == nil && node.rightChild == nil) {
		return true
	}
	var l, r int
	if node.leftChild != nil {
		l = node.leftChild.data
	}
	if node.rightChild != nil {
		r = node.rightChild.data
	}
	fmt.Printf("Node: %d, l:%d, r:%d\n", node.data, l, r)

	if node.data != l+r {
		return false
	}
	if node.leftChild != nil {
		lChildCheck := checkForChildreSum(node.leftChild)
		if !lChildCheck {
			return false
		}
	}
	if node.rightChild != nil {
		rChildCheck := checkForChildreSum(node.rightChild)
		if !rChildCheck {
			return false
		}
	}
	return true
}

func main() {
	// node structure
	//                         50
	//                     20      30
	//                   10 10    10 20
	// root node
	root := newNode(50)
	addNode(root, 50, 20, true)
	addNode(root, 50, 30, false)
	addNode(root, 20, 10, true)
	addNode(root, 20, 10, false)
	addNode(root, 30, 10, true)
	addNode(root, 30, 20, false)
	// inOrder(root)
	// fmt.Println(checkForChildreSum(root))
}
