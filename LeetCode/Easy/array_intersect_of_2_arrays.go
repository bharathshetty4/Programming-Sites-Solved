package main

/*
URL: https://leetcode.com/problems/intersection-of-two-arrays-ii/
Status: Success
Runtime: 6 ms, faster than 33.88% of Go online submissions for Intersection of Two Arrays II.
Memory Usage: 3 MB, less than 74.10% of Go online submissions for Intersection of Two Arrays II.
*/
import (
	"fmt"
)

func main() {
	fmt.Println("input: <>, output:", intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersect(nums1 []int, nums2 []int) []int {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 == 0 || len2 == 0 {
		return []int{}
	}
	if len1 < len2 {
		return process(nums1, nums2, len1, len2)
	}
	return process(nums2, nums1, len2, len1)
}

func process(mapArr, otherArr []int, mapLen, otherLen int) []int {
	numMap := make(map[int]int, mapLen)
	for i := 0; i < mapLen; i++ {
		numMap[mapArr[i]]++
	}
	totalNum := 0
	for i := 0; i < otherLen; i++ {
		if numMap[otherArr[i]] > 0 {
			if totalNum < mapLen {
				mapArr[totalNum] = otherArr[i]
			} else {
				mapArr = append(mapArr, otherArr[i])
			}
			numMap[otherArr[i]]--
			totalNum++
		}
	}
	if totalNum < mapLen {
		return mapArr[:totalNum]
	}
	return mapArr
}
