package main

/*
URL: https://leetcode.com/problems/reverse-integer/
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
Memory Usage: 2.1 MB, less than 80.67% of Go online submissions for Reverse Integer.
*/
import (
	"fmt"
)


func reverse(x int) int {
    MAX_INT := 2147483648-1
    isNeg := false
    if x < 0{
        isNeg = true
        x = x* -1
    }
    newVal := 0
    for {
        if x <= 0 {
            break
        }
        newVal = (newVal*10)+((x %10))
        if newVal >= MAX_INT {
            return 0
        }
        x = x/10
    }
    if isNeg{
        return newVal * -1
    }
    return newVal
}
