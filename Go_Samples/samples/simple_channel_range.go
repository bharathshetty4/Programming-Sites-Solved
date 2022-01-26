package main

import "fmt"

// This range iterates over each element as it’s received from queue.
// Because we closed the channel above, the iteration terminates after receiving the 2 elements.

func main() {
	queue := make(chan string, 2) // we can use unbuffered channel aswell with go routine
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

// This example also showed that it’s possible to close a non-empty channel
// but still have the remaining values be received.


// example for using unbuffered channel
func unbuffered() {
	done := make(chan string)
	go func() {
		for _, word := range []string{"foo", "bar", "baz"} {
			done <- word
		}
		close(done)
	}()
	for word := range done {
		fmt.Println(word)
	}
}
