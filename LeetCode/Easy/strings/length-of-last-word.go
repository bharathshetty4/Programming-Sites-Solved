package main

/*
URL: https://leetcode.com/problems/length-of-last-word/submissions/
Status: Accepted
Runtime: 0 ms, faster than 100.00% of Go online submissions for Length of Last Word.
Memory Usage: 2.1 MB, less than 28.84% of Go online submissions for Length of Last Word.
*/

func lengthOfLastWord(s string) int {
    s = strings.TrimSpace(s)
    fields := strings.Fields(s)
    return len(fields[len(fields)-1])
}
