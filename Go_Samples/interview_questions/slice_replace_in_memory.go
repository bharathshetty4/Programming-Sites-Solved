package main

import (
	"fmt"
)

// PROBLEM: For the given slice, check if the first cahracter is 'a' and remove the string from slice if so.

func main() {
	inputSlice := []string{"abc", "bbb", "aoi", "ccc", "dce", "abb"}
	prefixChar := "a"
	for i := len(inputSlice) - 1; i >= 0; i-- {
		currChar := inputSlice[i]
		if prefixChar == string(currChar[0]) {
			inputSlice = append(inputSlice[:i], inputSlice[i+1:]...)
		}
	}
	fmt.Println("input: <>, output:", inputSlice)
}
