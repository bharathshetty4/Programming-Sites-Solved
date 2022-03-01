package main

/*
URL: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
Status: Success
Runtime: 8 ms, faster than 88.29% of Go online submissions for Remove Duplicates from Sorted Array.
Memory Usage: 4.3 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted Array.
*/
import (
	"fmt"
)

func main() {
	fmt.Println("input: <>, output:", removeDuplicates([]int{1, 1, 2}))
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	cnt := 0
	for idx := 1; idx < len(nums); idx++ {
		if nums[cnt] == nums[idx] {
			continue
		}
		cnt++
		nums[cnt] = nums[idx]
	}
	return cnt + 1
}
