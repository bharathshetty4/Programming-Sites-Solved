//https://www.hackerrank.com/challenges/greedy-florist/problem?h_l=interview&playlist_slugs%5B%5D%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D%5B%5D=greedy-algorithms
// 9/12 failed
package main

import (
	"fmt"
)

// Complete the getMinimumCost function below.
func getMinimumCost(k int32, c []int32) int32 {
	cost := int32(0)
	iterNum := int32(0)
	flwLen := int32(len(c))
	for {
		if flwLen >= k {
			sum := int32(0)
			for i := flwLen - 1; i >= flwLen-k; i-- {
				sum = sum + c[i]
			}
			cost = cost + (iterNum+1)*sum
		}
		iterNum++
		flwLen = flwLen - k
		if flwLen < k {
			break
		}
	}
	for i := flwLen - 1; i >= 0; i-- {
		cost = cost + (iterNum+1)*c[i]
	}

	return cost
}

func main() {
	k := int32(2)
	c := []int32{2, 5, 6}
	res := getMinimumCost(k, c)
	fmt.Println(res)
}
