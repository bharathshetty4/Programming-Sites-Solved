package main

import "fmt"

func twoNums(y int, nums chan int) {
	nums <- y + 1
	//close(nums)
	nums <- y - 1
}

func main() {
	//single channel
	nums := make(chan int, 2)
	twoNums(5, nums)
	fmt.Println(<-nums)
	fmt.Println(<-nums)
}
