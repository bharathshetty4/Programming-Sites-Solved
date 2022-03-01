package main

/*
URL: https://www.geeksforgeeks.org/find-the-missing-number/
*/
import (
	"fmt"
)

// we are dealing with n-1 elements. and we need to calculate summation for n * n+1 /2
// array has to be continuos and the extra element should always be 1 to n+1
// input 1,3,4. Here 2 is missing. Here 'n' is 4.
func missingNumber(nums []int) int {
	len := len(nums)
	summation := ((len + 1) * ((len + 2) / 2))
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return summation - sum
}

// for the given array [2,3], calculate xor for idx and numbers with given 'n'. i.e 1^2 ^2^3 ^3
func missingNumberXor(nums []int) int {
	xor := len(nums) + 1
	for idx, num := range nums {
		xor = xor ^ num ^ idx + 1
	}
	return xor
}

func main() {
	fmt.Println("input: <>, output:", missingNumberXor([]int{2, 3}))
}
