package main

/*
URL: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
Status: Success
Runtime: 22 ms, faster than 27.24% of Go online submissions for Two Sum II - Input Array Is Sorted.
Memory Usage: 5.5 MB, less than 50.03% of Go online submissions for Two Sum II - Input Array Is 
*/
import (
	"fmt"
)


func twoSum(numbers []int, target int) []int {
    i := 0
    j:= len(numbers)-1
    for {
        if i >= j {
            return []int{}
        }
        sum := numbers[i]+numbers[j]
        if sum == target {
            return []int{i+1,j+1}
        }
        if sum < target {
            i++
        }else{
            j--
        }
    }
    return []int{}
}
