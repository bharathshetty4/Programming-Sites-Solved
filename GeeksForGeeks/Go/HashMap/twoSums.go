//https://leetcode.com/problems/two-sum/solution/
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for i,val := range nums{
		if idx,ok := sumMap[target-val];ok{
			return []int{idx,i}
		}
		sumMap[val] = i
	}
	return nil
}

func main(){
	values := []int{2, 7, 11, 15}
	target := 9
	idxs :=twoSum(values,target)
	fmt.Println(idxs)
	values = []int{20, 70, 1, 15}
	target = 38
	idxs =twoSum(values,target)
	fmt.Println(idxs)
}
