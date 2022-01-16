/*
Link: https://leetcode.com/problems/decompress-run-length-encoded-list
Runtime: 4 ms, faster than 93.56% of Go online submissions for Decompress Run-Length Encoded List.
Memory Usage: 5.9 MB, less than 51.98% of Go online submissions for Decompress Run-Length Encoded List.
*/
package main

func decompressRLElist(nums []int) []int {
	var val []int
	for i := 0; i < len(nums); i = i + 2 {
		for j := 0; j < nums[i]; j++ {
			val = append(val, nums[i+1])
		}
	}
	return val
}
