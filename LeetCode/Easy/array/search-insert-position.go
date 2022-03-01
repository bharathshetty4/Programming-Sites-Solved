package main

/*
URL: https://leetcode.com/problems/search-insert-position/
Status: Passed
Runtime: 5 ms, faster than 28.71% of Go online submissions for Search Insert Position.
Memory Usage: 3.3 MB, less than 23.78% of Go online submissions for Search Insert Position.
*/
import (
	"fmt"
)

// Find the right position to insert an element in the sorted array
func searchInsert(nums []int, target int) int {
	pos := 0

	numLen := len(nums)
	minIdx := 0
	maxIdx := numLen - 1
	for {
		midIdx := (maxIdx + minIdx) / 2
		fmt.Println("minIdx, midIdx, maxIdx, pos", minIdx, midIdx, maxIdx, pos)
		if minIdx > midIdx {
			break
		}
		if nums[midIdx] == target {
			return midIdx
		}
		if target > nums[midIdx] {
			if midIdx >= numLen {
				return numLen
			}
			pos = midIdx + 1
			minIdx = midIdx + 1
		} else {
			if midIdx <= 0 {
				return 0
			}
			pos = midIdx
			maxIdx = midIdx - 1
		}
	}
	return pos
}

func main() {
	fmt.Println("input: <>, output:", searchInsert([]int{1, 3, 5, 6, 45}, 2))
}
