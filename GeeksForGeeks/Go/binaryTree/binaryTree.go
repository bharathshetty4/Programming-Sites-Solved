package main

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

func main() {
	// node structure
	//                         10
	//                     20      30
	//                   40 50    60 70
	// root node
	root := newNode(10)
	addNode(root, 10, 20, true)
	addNode(root, 10, 30, false)
	addNode(root, 20, 40, true)
	addNode(root, 20, 50, false)
	addNode(root, 30, 60, true)
	addNode(root, 30, 70, false)
}
