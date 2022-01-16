/*
Link: https://leetcode.com/problems/valid-parentheses
91/ 91 test cases passed
Runtime: 0 ms
Memory Usage: 2.2 MB
*/

package main

import "fmt"

func canBeIncreasing(nums []int) bool {
	l := len(nums)
	if l <= 2 {
		return true
	}
	// lastDig := nums[0]
	// dupVal := 0
	// var numLess bool
	lastDig := nums[0]
	alreadyDec := false
	for i := 1; i < l; i++ {
		fmt.Printf("i %d\n", nums[i])

		if nums[i] <= nums[i-1] {
			if alreadyDec {
				fmt.Printf("i %d, compare with %d\n", nums[i], lastDig)
				return false
			}
			alreadyDec = true
			// check whether the last digit after omiting the decreasing num is lesses than nums[i]
			if lastDig > nums[i] {
				fmt.Printf("false as dup, lastdig %d nums[i] %d\n", lastDig, nums[i])
				return false
			}
			// lastDig = nums[i-1]
			// i--
			continue
		} else {
			lastDig = nums[i-1]
		}
	}
	return true
}

func main() {
	// fmt.Println(canBeIncreasing([]int{1, 2, 10, 5, 7}))
	// fmt.Println(canBeIncreasing([]int{2, 3, 1, 2}))
	// fmt.Println(canBeIncreasing([]int{1, 1}))         //true
	fmt.Println(canBeIncreasing([]int{100, 21, 100})) //true
}
