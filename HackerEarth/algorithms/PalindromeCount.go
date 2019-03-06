//https://www.hackerearth.com/practice/algorithms/dynamic-programming/2-dimensional/practice-problems/algorithm/palindrome-count-1/
package main

import (
	"fmt"
)

func isPal(str string) bool {
	n := len(str)
	for i := 0; i < n/2; i++ {
		if string(str[i]) != string(str[n-1-i]){
			return false
		}
	}
	return true
}

func countPalindromes(str string, n int) int {
	count := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if string(str[i]) == string(str[j]) {
				if isPal(str[i : j+1]) {
					count++
				}
			}
		}
	}
	return count
}
func main() {
	s := "dskjkd"
	n := len(s)
	c := countPalindromes(s, n)
	fmt.Print(n + c)
}
