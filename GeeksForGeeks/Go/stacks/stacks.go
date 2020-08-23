package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	items []int
	lock  sync.RWMutex
}

func (s *Stack) push(item int){
	s.lock.Lock()
	s.items = append(s.items,item)
	s.lock.Unlock()
}

func (s *Stack) pop(){
	s.lock.Lock()
	s.items = s.items[0:len(s.items)-1]
	s.lock.Unlock()
}

func main(){
	stackElem := new(Stack)
	stackElem.push(1)
	stackElem.push(5)
	stackElem.push(3)
	fmt.Println(stackElem.items)
	stackElem.pop()
	fmt.Println(stackElem.items)
	stackElem.push(4)
	fmt.Println(stackElem.items)
}