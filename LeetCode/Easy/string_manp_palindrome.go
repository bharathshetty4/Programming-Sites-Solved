package main

/*
URL: https://leetcode.com/problems/palindrome-number/
Status: Accepted
Runtime: 8 ms, faster than 95.51% of Go online submissions for Palindrome Number.
Memory Usage: 5.4 MB, less than 42.49% of Go online submissions for Palindrome Number.
*/
import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	t := strconv.Itoa(x)
	strLen := len(t)
	for i := 0; i < strLen/2; i++ {
		if t[i] != t[strLen-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("op: ", isPalindrome(-121))
	fmt.Println("op: ", isPalindrome(121))
}
