package main

/*
URL: https://leetcode.com/problems/valid-anagram
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Anagram.
Memory Usage: 2.7 MB, less than 88.55% of Go online submissions for Valid Anagram.
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
        alphaArr[s[i]-97]++
        alphaArr[t[i]-97]--
    }
    for i:=0;i<26;i++{
        if alphaArr[i] != 0 {
            return false
        }
    }
    return true
}
