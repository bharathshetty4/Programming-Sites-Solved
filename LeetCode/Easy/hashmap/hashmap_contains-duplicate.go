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

// Use struct instead of bool to reduce memory usage (along with num size for map)
/*
Runtime: 12 ms, faster than 99.94% of Go online submissions for Contains Duplicate.
Memory Usage: 7.9 MB, less than 77.92% of Go online submissions for Contains Duplicate.
*/
func containsDuplicateStruct(nums []int) bool {
    identity := make(map[int]struct{})
    for _, num:= range nums{
        if _, ok := identity[num]; ok {
            return true
        }
        identity[num] = struct{}{}
    }
    return false
}
