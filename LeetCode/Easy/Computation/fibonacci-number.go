package main

/*
URL: https://leetcode.com/problems/fibonacci-number/
Runtime: 21 ms, faster than 5.13% of Go online submissions for Fibonacci Number.
Memory Usage: 1.9 MB, less than 24.39% of Go online submissions for Fibonacci Number.
*/

import (
	"fmt"
)

func fib(n int) int {
    if n <=0 {
        return 0
    }
    if n <=1 {
        return 1
    }
    return fib(n-1)+fib(n-2)
}
