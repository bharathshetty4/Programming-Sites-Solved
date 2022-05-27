package main

/*
URL: https://leetcode.com/problems/counting-words-with-a-given-prefix
Runtime: 8 ms, faster than 20.43% of Go online submissions for Counting Words With a Given Prefix.
Memory Usage: 3.7 MB, less than 6.45% of Go online submissions for Counting Words With a Given Prefix.
*/

import (
	"fmt"
)


func prefixCount(words []string, pref string) int {
    cnt := 0
    prefLen := len(pref)
    for _, word := range words {
        if len(word) < prefLen {
            continue
        }
        if strings.HasPrefix(word,pref) {
            cnt++
        }
    }
    return cnt
}
