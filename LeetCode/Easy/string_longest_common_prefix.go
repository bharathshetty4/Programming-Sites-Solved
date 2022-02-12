package main

/*
URL: https://leetcode.com/problems/longest-common-prefix/
Runtime: 4 ms, faster than 31.22% of Go online submissions for Longest Common Prefix.
Memory Usage: 2.3 MB, less than 87.17% of Go online submissions for Longest Common Prefix.
*/
import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	index := 0
	for {
		breakOut := false
		if index >= len(strs[0]) {
			break
		}
		currLetter := strs[0][index] //current digit needed to be checked
		for _, s := range strs {
			if index >= len(s) || s[index] != currLetter {
				breakOut = true
				break
			}
		}
		if breakOut {
			break
		}
		index++
	}
	if index == 0 {
		return ""
	}
	return strs[0][0:index]
}
