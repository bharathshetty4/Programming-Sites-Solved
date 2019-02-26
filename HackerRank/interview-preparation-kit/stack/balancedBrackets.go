// https://www.hackerrank.com/challenges/balanced-brackets/problem?h_l=interview&playlist_slugs%5B%5D%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D%5B%5D=stacks-queues
// 3/18 is failing
package main

import "fmt"

type Stack []string

func (s Stack) ISEmpty() bool { return len(s) == 0 }

func (s *Stack) Push(v string) {
	(*s) = append((*s), v)
}

func (s *Stack) Pop() string {
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}

// Complete the isBalanced function below.
func isBalanced(s string) string {
	stack := Stack{}
	for _,val := range (s){
		v := string(val)
		switch v {
		case "(","{","[":
			stack.Push(v)
		case "}":
			if stack.ISEmpty() { return "NO"}
			poppedVal := stack.Pop()
			if poppedVal != "{" {
				return "NO"
			}
		case "]":
			if stack.ISEmpty() { return "NO"}
			poppedVal := stack.Pop()
			if poppedVal != "[" {
				return "NO"
			}
		case ")":
			if stack.ISEmpty() { return "NO"}
			poppedVal := stack.Pop()
			if poppedVal != "(" {
				return "NO"
			}
		}
	}
	return "YES"
}

func main() {
	s := "}][}}(}][))]"
	fmt.Println(isBalanced(s))
}
