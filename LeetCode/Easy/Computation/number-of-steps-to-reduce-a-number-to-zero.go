package main

/*
URL: https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero
Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of Steps to Reduce a Number to Zero.
Memory Usage: 1.9 MB, less than 83.90% of Go online submissions for Number of Steps to Reduce a Number to Zero.
*/

import (
	"fmt"
)

func numberOfSteps(num int) int {
    cnt := 0
    for {
        if num <= 0 {
            return cnt
        }
        cnt++
        if num %2 == 0 {
            num = num/2
        }else{
            num--
        }
    }
    return cnt
}
