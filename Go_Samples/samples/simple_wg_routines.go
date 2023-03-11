package main

import (
	"fmt"
	"sync"
)
// good read: https://about.sourcegraph.com/blog/building-conc-better-structured-concurrency-for-go
func runner1(wg *sync.WaitGroup) {
	defer wg.Done() // This decreases counter by 1
	fmt.Print("\nI am first runner")

}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am second runner")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go runner1(&wg)
	go runner2(&wg)
	wg.Wait()
}
