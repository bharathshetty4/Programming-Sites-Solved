package main

import (
	"fmt"
)

// PROBLEM: pair the numbers in the given sorted array.

func main() {
	inputNums := []int{1, 1, 3, 3, 4, 4, 5, 6, 8, 11, 12, 14, 15, 16, 16}
	// print 1, 3-6, 8, 11-12, 14-16

	for i := 0; i < len(inputNums); i++ {
		startElem := inputNums[i]
		endElem := inputNums[i]
		for {
			if i >= len(inputNums)-1 {
				break
			}
			if inputNums[i] == inputNums[i+1] {
				i++
			} else if endElem+1 == inputNums[i+1] {
				endElem++
				i++
			} else {
				break
			}
		}
		if startElem != endElem {
			fmt.Printf("%d-%d,", startElem, endElem)
		} else {
			fmt.Printf("%d,", startElem)
		}
	}

}
