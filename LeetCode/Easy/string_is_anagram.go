package main

/*
URL: https://leetcode.com/problems/valid-anagram
Runtime: 5 ms, faster than 69.71% of Go online submissions for Valid Anagram.
Memory Usage: 2.8 MB, less than 93.73% of Go online submissions for Valid Anagram.
*/

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    alphaArr := make([]int,26)
    for i:=0;i<len(s);i++{
        alphaArr[s[i]-'a']++
        alphaArr[t[i]-'a']--
    }
    for i:=0;i<26;i++{
        if alphaArr[i] != 0 {
            return false
        }
    }
    return true
}
