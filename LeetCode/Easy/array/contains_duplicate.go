package main

/*
URL: https://leetcode.com/problems/contains-duplicate
Status: Success
Runtime: 96 ms, faster than 56.49% of Go online submissions for Contains Duplicate.
Memory Usage: 9.1 MB, less than 48.28% of Go online submissions for Contains Duplicate.
*/
import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	identity := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := identity[num]; ok {
			return true
		}
		identity[num] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println("input: <>, output:", containsDuplicate([]int{1, 1, 2}))
}
