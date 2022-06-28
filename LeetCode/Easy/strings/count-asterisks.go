package main

/*
URL: https://leetcode.com/problems/count-asterisks/submissions/
Status: Accepted
Runtime: 2 ms, faster than 49.12% of Go online submissions for Count Asterisks.
Memory Usage: 2.1 MB, less than 52.63% of Go online submissions for Count Asterisks.
*/

func countAsterisks(s string) int {
    totalAstr := 0
    strs := strings.Split(s,"|")
    // we need to get the asteriks count of even indexes as odd index strings
    // belong to the pair
    for i:=0;i<len(strs);i=i+2 {
        totalAstr+=numAsterisk(strs[i])
    }
    return totalAstr
}

func numAsterisk(s string) int {
    return len(strings.Split(s,"*"))-1
}
