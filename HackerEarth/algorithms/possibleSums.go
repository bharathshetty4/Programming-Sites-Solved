package main

import (
	"fmt"
)

func getPossSums(nums []int) int {
	count := 0
	uniqueNums := make(map[int]bool)
	uniqueNums[0] = true
	n := len(nums)
	for i := 0; i < n; i++ {
		if uniqueNums[nums[i]] == false {
			uniqueNums[nums[i]] = true
			count++
		}
		for j := i + 1; j < n; j++ {
			if uniqueNums[nums[i]+nums[j]] == false {
				uniqueNums[nums[i]+nums[j]] = true
				count++
			}
		}
	}

	fmt.Println(uniqueNums)
	return count + 1
}

func main() {
	nums := []int{1, 2, 3}
	sums := getPossSums(nums)
	fmt.Println(sums)
}
