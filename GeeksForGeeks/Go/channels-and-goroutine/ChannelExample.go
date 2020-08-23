package main

import (
	"fmt"
)

func sumTwo(x, y int, sum chan int) {
	sum <- x + y
	//if x%2 != 0 {
	//	time.Sleep(time.Second)
	//}
}

func main() {
	//single channel
	sum := make(chan int)
	go func() {
		sumTwo(5, 10, sum)

	}()
	fmt.Println(<-sum)
	go func() {
		sumTwo(10, 15, sum)
	}()
	fmt.Println(<-sum)
}
