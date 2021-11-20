/*
Link: https://leetcode.com/problems/two-sum

29 / 29 test cases passed
Runtime: 4 ms
Memory Usage: 3.7 MB
*/
package main

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for i, val := range nums {
		if idx, ok := sumMap[target-val]; ok {
			return []int{idx, i}
		}
		sumMap[val] = i
	}
	return nil
}
