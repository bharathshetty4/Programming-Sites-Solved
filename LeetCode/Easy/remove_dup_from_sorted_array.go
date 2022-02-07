package main

/*
URL: https://leetcode.com/problems/remove-duplicates-from-sorted-array
Status: Success
Runtime: 5 ms, faster than 91.82% of Go online submissions for Remove Duplicates from Sorted Array.
Memory Usage: 4.4 MB, less than 58.37% of Go online submissions for Remove Duplicates from Sorted Array.
*/
import (
	"fmt"
)

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

func main() {
	fmt.Println("input: <>, output:", removeDuplicates([]int{1, 1, 2}))
}
