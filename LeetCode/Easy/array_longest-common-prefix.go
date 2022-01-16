package main

/*
URL: https://leetcode.com/problems/longest-common-prefix
Status: Success
Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Longest Common Prefix.
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

func main() {
	fmt.Println("input: <>, output:", longestCommonPrefix([]string{"car", "ca"}))
}
