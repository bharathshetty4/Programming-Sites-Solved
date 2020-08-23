//https://www.geeksforgeeks.org/next-greater-element/

package main

import (
	"fmt"
)

func nextGreaterElement(elements []int) {
	eleLen := len(elements)
	nextBiggerElements := make([]int, eleLen)
	for i := eleLen - 1; i >= 0; i-- {
		nextMinimum := -1
		for j := i + 1; j <= eleLen-1; j++ {
			if elements[j] > elements[i] {
				if nextMinimum == -1 || nextMinimum > elements[j] {
					nextMinimum = elements[j]
				}
			}
		}
		nextBiggerElements[i] = nextMinimum
	}
	fmt.Printf("Element is %v\n", nextBiggerElements)
}

func main() {
	ele := []int{4, 5, 2, 25}
	nextGreaterElement(ele)

	ele = []int{13, 7, 6, 12}
	nextGreaterElement(ele)

	ele = []int{11, 13, 21, 3}
	nextGreaterElement(ele)

	ele = []int{2,5,9,6,3,4,8,15,12}
	nextGreaterElement(ele)
}
