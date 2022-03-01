package main

/*
URL: https://leetcode.com/problems/missing-number/submissions/
Status: Success
Runtime: 20 ms, faster than 67.82% of Go online submissions for Missing Number.
Memory Usage: 7.4 MB, less than 5.24% of Go online submissions for Missing Number.
*/
import (
	"fmt"
)

// for the given number from 0 to n, find the missing number in (n-1) array
// ex: input: [0,1] output: 2. 0 to n is 0,1,2 and 0,1 is part of the array.
func missingNumber(nums []int) int {
    xor := len(nums)
	for idx, num := range nums {
		xor = xor ^ num ^ idx
	}
    return xor
}
