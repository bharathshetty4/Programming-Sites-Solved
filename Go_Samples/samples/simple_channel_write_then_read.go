package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 4; i++ {
		go updateChannel(i, ch)
	}

	for i := 0; i < 2; i++ {
		time.Sleep(time.Second)
		recvdVal := <-ch
		fmt.Println("Reading data", recvdVal)
	}
	close(ch)
}

func updateChannel(val int, ch chan int) {
	fmt.Println("writing data", val)
	ch <- val
}
/*
output:
writing data 3
writing data 2
writing data 1
writing data 0
Reading data 3
Reading data 2
*/
