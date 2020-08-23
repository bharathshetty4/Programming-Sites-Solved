package main

import (
	"fmt"
	"sync"
)

func sumIt(x, y int, sum chan int) {
	sum <- x + y

	//time.Sleep(time.Second)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	//single channel
	sum := make(chan int)
	go func() {
		go sumIt(5, 10, sum)
		fmt.Println(<-sum)
		wg.Done()
	}()

	go func() {
		go sumIt(10, 15, sum)
		fmt.Println(<-sum)
		wg.Done()
	}()
	wg.Wait()
}
