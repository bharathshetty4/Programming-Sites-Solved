package main

/*
URL: https://leetcode.com/problems/subsets
Status: Accepted
Runtime: 0 ms, faster than 100.00% of Go online submissions for Subsets.
Memory Usage: 2.6 MB, less than 19.72% of Go online submissions for Subsets.
*/
import (
	"fmt"
)

func subsets(nums []int) [][]int {
	var output []int
	var sets [][]int
	getSubsets(nums, output, 0, &sets)
	return sets
}

func getSubsets(nums, output []int, index int, sets *[][]int) {
	if index >= len(nums) {
		// fmt.Printf("nums %#v, output %#v, index %d, sets %#v\n",nums,output,index, sets)
		// need a local output as the output gets updated in golang and mess up the list if not used
		var localOutput []int
		for _, val := range output {
			localOutput = append(localOutput, val)
		}
		// fmt.Printf("localOutput %#v\n",localOutput)
		*sets = append(*sets, localOutput)
		return
	}
	// exclude
	getSubsets(nums, output, index+1, sets)

	// include
	output = append(output, nums[index])
	getSubsets(nums, output, index+1, sets)

}

func main() {
	fmt.Println("input: {1,2,3}, output:", subsets([]int{1, 2, 3}))
}
