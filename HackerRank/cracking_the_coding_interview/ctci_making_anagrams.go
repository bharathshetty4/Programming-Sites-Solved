// https://www.hackerrank.com/challenges/ctci-making-anagrams/problem
package main

import (
	"fmt"
	"math"
)

func main() {
	var firstString, secondString string
	fmt.Scanf("%s\n%s", &firstString, &secondString)

	var letterCount = make(map[string]int)
	var totalCount float64
	for _, letter := range firstString {
		letterCount[string(letter)] = letterCount[string(letter)] + 1
	}
	for _, letter := range secondString {
		letterCount[string(letter)] = letterCount[string(letter)] - 1
	}

	for _, count := range letterCount {
		totalCount = totalCount + math.Abs(float64(count))
	}
	fmt.Printf("%d", int(totalCount))
}
