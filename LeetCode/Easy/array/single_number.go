package main

/*
URL: https://leetcode.com/problems/single-number
Status: Success
Runtime: 14 ms, faster than 90.71% of Go online submissions for Single Number.
Memory Usage: 6.3 MB, less than 73.26% of Go online submissions for Single Number.
*/
import (
	"fmt"
)

func main() {
	fmt.Println("input: <>, output:", singleNumber([]int{2, 2, 1, 3, 1}))
}

func singleNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum = sum ^ num
	}
	return sum
}
