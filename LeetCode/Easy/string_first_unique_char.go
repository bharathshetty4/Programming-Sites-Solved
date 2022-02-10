package main

/*
URL: https://leetcode.com/problems/contains-duplicate
Status: Success
Runtime: 108 ms, faster than 16.49% of Go online submissions for Contains Duplicate.
Memory Usage: 9.7 MB, less than 48.28% of Go online submissions for Contains Duplicate.
*/
import (
	"fmt"
)


// mycode 
func firstUniqChar(s string) int {
    runes := []rune(s)
    strMap := make(map[string]int,len(s))
    for _, r:= range runes{
        strMap[string(r)]++
    }
    for idx, r:= range runes{
        if strMap[string(r)] <= 1{
            return idx
        }
    }
    return -1
}

// best sol
func firstUniqChar(s string) int {
    var counts [26]int
    for i := 0; i < len(s); i++ {
        counts[s[i] - 'a']++
    }
    
    for i := 0; i < len(s); i++ {
        if counts[s[i] - 'a'] == 1 {
            return i
        }
    }
    
    return -1
}
