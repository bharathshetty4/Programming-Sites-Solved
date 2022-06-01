package main

/*
URL: https://leetcode.com/problems/intersection-of-two-arrays-ii/
Status: Success
Runtime: 7 ms, faster than 13.10% of Go online submissions for Running Sum of 1d Array.
Memory Usage: 2.7 MB, less than 85.71% of Go online submissions for Running Sum of 1d Array.
*/
import (
	"fmt"
)

func runningSum(nums []int) []int {
    if len(nums) <=1{
        return nums
    }
    for i:=1;i<len(nums);i++{
        nums[i]+=nums[i-1]
    }
    return nums
}
