/*
Link: https://leetcode.com/problems/contains-duplicate
Runtime: 16 ms, faster than 94.94% of Go online submissions for Contains Duplicate.
Memory Usage: 8.9 MB, less than 31.92% of Go online submissions for Contains Duplicate.
*/
package main

func containsDuplicate(nums []int) bool {
    identity := make(map[int]bool)
    for _, num:= range nums{
        if _, ok := identity[num]; ok {
            return true
        }
        identity[num] = true
    }
    return false
}
