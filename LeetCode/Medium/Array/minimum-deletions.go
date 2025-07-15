package main

/*
URL: https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/
Status: Success
Runtime: 13 ms, faster than 98.39% of Go online submissions for Minimum Deletions to Make Character Frequencies Unique.
Memory Usage: 6.5 MB, less than 85.48% of Go online submissions for Minimum Deletions to Make Character Frequencies Unique.
*/

import (
	"fmt"
)


func minDeletions(s string) int {
    var charCnt [26]int
    // get count of each character
    for _, strChar := range s{
        charCnt[strChar-'a']++
    }
    fmt.Println(charCnt)
    delCnt :=0 
    cntMap := make(map[int]struct{})
    // for each character, if the count is already used by any other character
    // loop until the count is unique or zero
    for i:=0;i<26;i++{
        cntVal := charCnt[i]
        for cntVal>0{
            if _,ok:=cntMap[cntVal];ok{
                delCnt++
                cntVal--
            }else{
                cntMap[cntVal]=struct{}{}
                break
            }
        }
    }
    return delCnt
}
