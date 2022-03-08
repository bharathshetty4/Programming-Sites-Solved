package main

/*
URL: https://leetcode.com/problems/reverse-words-in-a-string/
Status: Accepted
Runtime: 3 ms, faster than 61.19% of Go online submissions for Reverse Words in a String.
Memory Usage: 2.9 MB, less than 82.09% of Go online submissions for Reverse Words in a String.
*/

import (
	"strings"
)

func reverseWords(s string) string {
    if s == ""{
        return ""
    }
    f := strings.Fields(s)
    for i:=0;i<len(f)/2;i++{
        f[i],f[len(f)-1-i] = f[len(f)-1-i],f[i]
    }
    return strings.Join(f," ")
}
