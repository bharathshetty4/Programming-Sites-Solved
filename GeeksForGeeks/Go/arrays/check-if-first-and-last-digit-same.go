package main

import (
	"fmt"
)

func isFirstLastDigitSame(arr []int) {
	for i, val := range arr {
		first := val % 10
		for {
			if val < 10 {
				break
			}
			val = val / 10
		}
		if first == val {
			fmt.Println(arr[i])
		}
	}
}

func main() {
	arr := []int{111, 123, 343}
	isFirstLastDigitSame(arr)
}
