
package main

/*
URL: https://leetcode.com/problems/maximum-erasure-value/
Status: Success
Runtime: 1149 ms, faster than 8.00% of Go online submissions for Maximum Erasure Value.
Memory Usage: 9.3 MB, less than 56.00% of Go online submissions for Maximum Erasure Value.
*/

import (
	"fmt"
)


// modified the https://leetcode.com/problems/longest-substring-without-repeating-characters,so varibale name might not make sense
func maximumUniqueSubarray(nums []int) int {
    if len(nums) == 1 {
		return nums[0]
	}
	longestSubStr := 0
	startSubStrIndex := 0
	endSubStrIndex := 0
	charIndex := make(map[int]int)
	for idx, char := range nums {
		// a duplicate character found.
		if charOldIndex, ok := charIndex[char]; ok {
			if charOldIndex < startSubStrIndex {
				// if the duplicate char found exist before the start index, nothing to do and continue the computation
			} else {
                sum := getSumBetweenIndex(nums, startSubStrIndex, endSubStrIndex)
				if longestSubStr < sum {
					longestSubStr = sum
				}
				startSubStrIndex = charOldIndex + 1
			}
		}
		// keep incrementing the end index of the substring
		endSubStrIndex++
		charIndex[char] = idx
	}

    lastSum := getSumBetweenIndex(nums, startSubStrIndex, endSubStrIndex)
    if longestSubStr < lastSum {
		longestSubStr = lastSum
	}
	return longestSubStr
}

func getSumBetweenIndex(nums []int, start, end int) int{
    sum := 0
    for i:=start;i<len(nums)&&i<end;i++{
        sum = sum +nums[i]
    }
    return sum
}
