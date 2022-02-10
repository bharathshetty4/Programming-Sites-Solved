package main

/*
URL: https://leetcode.com/problems/rotate-array
Status: Success
Runtime: 4790 ms, faster than 5.01% of Go online submissions for Rotate Array.
Memory Usage: 8.1 MB, less than 71.33% of Go online submissions for Rotate Array.
*/
import (
	"fmt"
)

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}
	numLen := len(nums)
	// rotate two number at a time, we can take also take some definite number of slices such as gcd and move those at a time
	// instead of moving only 2 numbers at a time
	for j := 0; j < k; j++ {
		temp := nums[numLen-1]
		for i := 1; i < numLen; {
			nums[numLen-i] = nums[numLen-1-i]
			if numLen-2-i >= 0 {
				nums[numLen-i-1] = nums[numLen-2-i]
			}
			i = i + 2
		}
		nums[0] = temp
	}

}

func main() {
	fmt.Println("input: <>, output:")
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
