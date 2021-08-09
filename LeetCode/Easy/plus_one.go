package main

/*
URL: https://leetcode.com/problems/plus-one
Status: Accepted. 111 / 111 test cases passed.
Runtime: 0 ms, faster than 100.00% of Go online submissions for Plus One.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Plus One.
*/
import (
	"fmt"
)

func plusOne(digits []int) []int {
	lenOfDig := len(digits)
	extraDigit := false
	for {
		if lenOfDig == 0 {
			if extraDigit {
				newDig := []int{1}
				return append(newDig, digits...)
			}
			return digits
		}
		extraDigit = false
		if digits[lenOfDig-1] != 9 {
			digits[lenOfDig-1] = digits[lenOfDig-1] + 1
			return digits
		}
		extraDigit = true
		digits[lenOfDig-1] = 0
		lenOfDig--
	}

}
func main() {
	fmt.Println("input: [9], output:", plusOne([]int{9}))
}
