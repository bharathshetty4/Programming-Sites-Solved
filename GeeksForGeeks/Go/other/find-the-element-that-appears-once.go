// https://www.geeksforgeeks.org/find-element-appears-array-every-element-appears-twice
// https://www.geeksforgeeks.org/find-the-element-that-appears-once
package main

import (
    "fmt"
)
func getSingleAppeared(elements []int) int {
	sum := elements[0]
	for i := 1; i < len(elements); i++ {
		sum = sum ^ elements[i]
	}
	return sum
}

func main() {
	fmt.Println("[]string{1,2,1,3,2} :", getSingleAppeared([]int{1, 2, 1, 3, 2}))
    fmt.Println("[]string{1,3,1,3,2} :", getSingleAppeared([]int{1, 3, 1, 3, 2}))
}
