package main

/*
URL: https://leetcode.com/problems/remove-element
Status: Success
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
Memory Usage: 2.1 MB, less than 86.73% of Go online submissions for Remove Element.
*/
import (
	"fmt"
)

func removeElement(nums []int, val int) int {
    if len(nums) <0||(len(nums)==1 && nums[0]==val) {
		return 0
	}
	cnt := 0
	for idx := 0; idx < len(nums); idx++ {
		if val == nums[idx] {
			continue
		}
		nums[cnt] = nums[idx]
        cnt++
	}
	return cnt
}

func main() {
	fmt.Println("input: <>, output:",removeElement([]int{3,2,2,3},3))
}
