package main

/*
  URL: https://leetcode.com/submissions/detail/386202542/?from=/explore/challenge/card/30-day-leetcoding-challenge/528/week-1/3283/
  Status: Accepted

  16 / 16 test cases passed.
  Runtime: 12 ms
  Memory Usage: 5.5 MB
  Runtime: 12 ms, faster than 63.41% of Go online submissions for Add Two Numbers.
  Memory Usage: 5.5 MB, less than 28.37% of Go online submissions for Add Two Numbers.
*/
import (
	"fmt"
)

func singleNumber(nums []int) int {
	numCount := make(map[int]bool)
	numLen := len(nums)
	for i := 0; i < numLen/2; i++ {
		if nums[i] == nums[numLen-1-i] {
			continue
		}
		if numCount[nums[i]] {
			delete(numCount, nums[i])
		} else {
			numCount[nums[i]] = true
		}
		if numCount[nums[numLen-1-i]] {
			delete(numCount, nums[numLen-1-i])
		} else {
			numCount[nums[numLen-1-i]] = true
		}
	}

	if numLen%2 == 1 {
		if numCount[nums[numLen/2]] == false {
			return nums[numLen/2]
		}
		delete(numCount, nums[numLen/2])
	}
	for k := range numCount {
		return k
	}
	return 0
}

func main() {
	num := singleNumber([]int{1, 3, 1, -1, 3})
	fmt.Printf("number got %+v\n", num)
}
