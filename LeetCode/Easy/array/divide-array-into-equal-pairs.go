package main

/*
URL: https://leetcode.com/problems/divide-array-into-equal-pairs/
Status: Success
Runtime: 23 ms, faster than 34.57% of Go online submissions for Divide Array Into Equal Pairs.
Memory Usage: 5.8 MB, less than 40.74% of Go online submissions for Divide Array Into Equal Pairs.
*/
import (
	"fmt"
)

func main() {
	fmt.Println("input: <>, output:", intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}

func divideArray(nums []int) bool {
    numCount := make(map[int]int)
    for _, num := range nums {
        numCount[num]++
    }
    for _, val := range numCount{
        if val % 2 != 0 {
            return false
        }
    }
    return true
}
