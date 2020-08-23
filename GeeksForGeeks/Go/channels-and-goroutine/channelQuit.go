package main

import (
	"fmt"
	"math/rand"

	"time"
)

func quitChannel(quit chan bool) {
	time.Sleep(time.Second)
	// Do stuff
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rand := r1.Int63n(100)
	if rand%2 == 0 {
		quit <- true
		quit <- true
	}
	fmt.Println("got odd, closing channel", rand)
	close(quit)
}

func main() {
	quit := make(chan bool)
	go quitChannel(quit)
	for {
		select {
		case <-quit:
			fmt.Println("Received quit")
			return
		default:
			println("Waiting to get data or close channel...")
			time.Sleep(time.Millisecond * 300)
		}
	}
}
