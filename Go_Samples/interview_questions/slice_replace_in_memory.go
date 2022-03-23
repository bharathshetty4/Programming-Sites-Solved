package main

import (
	"fmt"
)

// PROBLEM: For the given slice, check if the first cahracter is 'a' and remove the string from slice if so.

func main() {
	inputSlice := []string{"abc", "bbb", "oi", "ccc", "dce", "abb"}
	prefixChar := "a"
	idxPtr := 0
	for i := 0; i < len(inputSlice); i++ {
		if prefixChar == string(inputSlice[i][0]) {
			continue
		}
		inputSlice[idxPtr] = inputSlice[i]
		idxPtr++
	}
	fmt.Println("input: <>, output:", inputSlice[:idxPtr])
}
