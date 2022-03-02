package main

/*
URL: https://leetcode.com/problems/contains-duplicate
URL: https://leetcode.com/problems/first-unique-character-in-a-string
Status: Success
Runtime: 10 ms, faster than 83.10% of Go online submissions for First Unique Character in a String.
Memory Usage: 6.4 MB, less than 19.25% of Go online submissions for First Unique Character in a String.
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
