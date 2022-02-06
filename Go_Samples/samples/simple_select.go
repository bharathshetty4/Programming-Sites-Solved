// Go program to illustrate the
// concept of select statement
package main

import (
	"fmt"
	"time"
)

func portal1(channel1 chan string) {
	channel1 <- "Welcome to channel 1"
}

func portal2(channel2 chan string) {
	channel2 <- "Welcome to channel 2"
}

// main function
func main() {

	// Creating channels
	// NOTE: make it buffered channel, so that when we cant read this channel, it will not end up leaking go routine
	R1 := make(chan string,1)
	R2 := make(chan string,1)

	// calling function 1 and
	// function 2 in goroutine
	go portal1(R1)
	go portal2(R2)
	i := 0
	// Go’s select lets you wait on multiple channel operations.
	// can have 'default:' case or timer inside.
	// NOTE: If a select statement has non-blocking case like default or timer, it can choose any one of those and here
	// in this case we may never hit the timer at all
	select {

	// case 1 for portal 1
	case op1 := <-R1:
		fmt.Println(op1)

	// case 2 for portal 2
	case op2 := <-R2:
		fmt.Println(op2)

	case <-time.After(time.Nanosecond):
		fmt.Printf("\nTimeout\n")
		return
	default:
		if i == 2 {
			break
		}
		fmt.Println("def")
	}

}
