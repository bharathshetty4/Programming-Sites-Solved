//https://www.geeksforgeeks.org/check-for-balanced-parentheses-in-an-expression/
package main

import (
	"fmt"
)

var MatchingBraces = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

func isBlanced(expression string) bool {
	fmt.Printf("Passed string is %s : ",expression)
	if len(expression)%2 != 0 {
		return false
	}
	m := len(expression) / 2
	for i := 0; i < m; i++ {
		if string(expression[m-1-i]) != MatchingBraces[string(expression[m+i])] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isBlanced("([])"))
	fmt.Println(isBlanced("([)"))
	fmt.Println(isBlanced("([[[[)"))
	return
}
