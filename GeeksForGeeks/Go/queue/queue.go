package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	items []int
	lock  sync.RWMutex
}

func (q *Queue) enqueue(item int){
	q.lock.Lock()
	q.items = append(q.items,item)
	q.lock.Unlock()
}

func (q *Queue) deque(){
	q.lock.Lock()
	q.items = q.items[1:len(q.items)]
	q.lock.Unlock()
}

func main(){
	queueElem := new(Queue)
	queueElem.enqueue(1)
	queueElem.enqueue(5)
	queueElem.enqueue(3)
	fmt.Println(queueElem.items)
	queueElem.deque()
	fmt.Println(queueElem.items)
	queueElem.enqueue(4)
	fmt.Println(queueElem.items)
}