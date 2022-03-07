package main

/*
URL: https://leetcode.com/problems/ransom-note
Status: Success
Runtime: 7 ms, faster than 80.00% of Go online submissions for Ransom Note.
Memory Usage: 3.9 MB, less than 84.13% of Go online submissions for Ransom Note.
*/

func canConstruct(ransomNote string, magazine string) bool {
	var letCnt [26]int
	for _, char := range magazine {
		letCnt[char-97]++
	}
	for _, char := range ransomNote {
		letCnt[char-97]--
		if letCnt[char-97] < 0 {
			return false
		}
	}
	return true
}
