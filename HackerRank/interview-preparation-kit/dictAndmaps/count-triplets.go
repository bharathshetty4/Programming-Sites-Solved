//https://www.hackerrank.com/challenges/count-triplets-1/problem?h_l=interview&playlist_slugs%5B%5D%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D%5B%5D=dictionaries-hashmaps
//5/13 is failing
package main

import (
	"fmt"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	var res int64
	count := int64(0)

	lastItem,nextItem := int64(1),int64(1)
	for i := 0; i < len(arr); i++ {
		if count == 0 && arr[i] != 1 {
			continue
		}
		if count == 0 && arr[i] == 1 {
			count ++
			nextItem = 1 * r
			lastItem = 1
			continue
		}
		if arr[i] == nextItem || arr[i]==lastItem {
			count++
			if arr[i] == nextItem {
				nextItem = nextItem * r
			}
			lastItem = arr[i]
		}
	}

	fmt.Println(count)
	if count == 3 {
		return 1
	}
	if count > 3 {
		res = 2 * (count - 3)
	}
	return res
}

func main() {
	nums := []int64{1, 2, 2, 4}
	r := int64(2)
	res := countTriplets(nums, r)
	fmt.Println(res)
}
