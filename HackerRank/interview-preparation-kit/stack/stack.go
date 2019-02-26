package main

import (
	"fmt"
)

type StackOld []int

func (s StackOld) ISEmpty() bool { return len(s) == 0 }

func (s *StackOld) Push(v int) {
	(*s) = append((*s), v)
}

func (s *StackOld) Pop() int {
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}

func main() {
	s := StackOld{}
	fmt.Println(s.ISEmpty())
	s.Push(4)
	s.Push(5)
	s.Push(6)
	fmt.Println("Before Stack", s)
	s.Pop()
	fmt.Println("After Stack", s)
}
