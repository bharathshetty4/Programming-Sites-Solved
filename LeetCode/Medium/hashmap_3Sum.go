package main

/*
URL: https://leetcode.com/problems/3sum
Status: Success. 318 / 318 test cases passed.
Runtime: 1084 ms, faster than 13.20% of Go online submissions for 3Sum.
Memory Usage: 9.7 MB, less than 13.30% of Go online submissions for 3Sum.
*/

import (
	"fmt"
	"time"
)

type numMapStruct struct {
	exist bool
	index int
}

func threeSum(nums []int) [][]int {
	var threeSums [][]int

	numMap := make(map[int]numMapStruct)
	numLen := len(nums)
	for index, val := range nums {
		// create a map of each numbers given in an array. Store the index aswell and a bool to indicate
		//that the number exist in the struct instead of checking it using checker (ok checker)
		numMap[val] = numMapStruct{
			exist: true,
			index: index,
		}
	}
	// loop from first digit to the last 2nd digit
	for i := 0; i < numLen-2; i++ {
		for j := i + 1; j < numLen-1; j++ {
			// if the 3rd digit which is needed to form a 3 sum found to be i or j, continue with the next number
			if numMap[(nums[i]+nums[j])*-1].index == i || numMap[(nums[i]+nums[j])*-1].index == j {
				continue
			}
			if numMap[(nums[i]+nums[j])*-1].exist {
				// got 3 sum numbers. Sort the numbers so that we dont end up having duplicate 3 sums
				threeSums = sortThreeSums(threeSums, ((nums[i] + nums[j]) * -1), nums[i], nums[j])
			}
		}
	}

	return threeSums
}

func sortThreeSums(threeSums [][]int, i, j, k int) [][]int {
	var thisNums []int
	if i >= j && i >= k {
		//i is greater
		if j >= k {
			thisNums = []int{k, j, i}
		} else {
			thisNums = []int{j, k, i}
		}

	} else if j >= i && j >= k {
		//j is greater
		if i >= k {
			thisNums = []int{k, i, j}
		} else {
			thisNums = []int{i, k, j}
		}

	} else {
		//k is greater
		if i >= j {
			thisNums = []int{j, i, k}
		} else {
			thisNums = []int{i, j, k}
		}
	}
	for _, sums := range threeSums {
		if thisNums[0] == sums[0] && thisNums[1] == sums[1] && thisNums[2] == sums[2] {
			return threeSums
		}
	}
	return append(threeSums, thisNums)
}

func main() {
	sTime := time.Now()
	fmt.Println("op:", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println("\n op:", threeSum([]int{1, 2, -2, -1}))
	fmt.Println("\n op:", threeSum([]int{3, -2, 1, 0}))
	elapsed := time.Since(sTime)
	fmt.Println("time taken:", elapsed)
}
