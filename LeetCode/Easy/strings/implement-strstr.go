package main

/*
URL: https://leetcode.com/problems/implement-strstr/
Status: Success
Runtime: 3 ms, faster than 54.86% of Go online submissions for Implement strStr().
Memory Usage: 2 MB, less than 78.94% of Go online submissions for Implement strStr().
*/
import (
	"fmt"
)

func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return -1
    }
    f := needle[0]
    hayLen := len(haystack)
    needleLen := len(needle)
    i := 0
    for {
        if i >= hayLen || hayLen-i < needleLen{
            return -1
        }
        if haystack[i]==f{
            // check all chars of needle
            hayIndex := i
            needleIdx :=0 
            for j:= hayIndex;j<hayLen&&needleIdx<needleLen;j++{
                if haystack[j]!=needle[needleIdx]{
                    break
                }
                needleIdx++
            }
            if needleIdx == needleLen{
                return hayIndex
            }
        }
        i++
    }
}
