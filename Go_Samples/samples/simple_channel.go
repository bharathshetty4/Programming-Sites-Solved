package main

import (
	"fmt"
	"time"
)

// Blogs to read:
// https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html

type chanStruct struct {
	intVal int
	strVal string
}

func main() {
	ch := make(chan chanStruct, 0) // unbuffered channel
	go updateChannel(ch)
	fmt.Println("Hello from channel:", <-ch) // waits until the data is written as it is unbuffered

	bufferedChannelSameThread()
	unbufferedWithDifferentThread()
	bufferedWithDifferentThread()
}

func updateChannel(ch chan chanStruct) {
	obj := chanStruct{intVal: 4, strVal: "bks"}
	ch <- obj
}

/*
output:
Hello from channel: {4 bks}
*/

func bufferedChannelSameThread() {
	// create a buffered channel, which can write as many times as needed without worrying about the write.
	queue := make(chan string, 2) // this has to be 2 if we are reading and writing in the same thread
	queue <- "one"
	queue <- "two"
	close(queue) // make sure that you close the channel before ranging over it. Otherwise for loop will never break.

	for elem := range queue {
		fmt.Println(elem)
	}
}

/*
output:
one
two
*/

// range for unbuffered channel by using multiple thread
func unbufferedWithDifferentThread() {
	done := make(chan string) // you can make this buffered to process faster, as 'range' will wait until the data is written on the channel after a 'read'
	go func() {
		for _, word := range []string{"foo", "bar", "baz"} {
			done <- word
		}
		close(done)
		//done <- "dd" NOT VALID. after closing the channel, you cannot send anything.
	}()
	for word := range done {
		fmt.Println(word)
	}
	fmt.Println("read after closing the channel:", <-done) // read after closing the channel. VALID.
}

/*
output:
foo
bar
baz
read after closing the channel:
*/

func bufferedWithDifferentThread() {
	ch := make(chan int, 1) // we can write multiple times if the channel is used in different thread even with the less size
	for i := 0; i < 4; i++ {
		go updateBufferedChannel(i, ch)
	}

	for i := 0; i < 2; i++ {
		time.Sleep(time.Second)
		recvdVal := <-ch
		fmt.Println("Reading data", recvdVal)
	}
	close(ch)
}

func updateBufferedChannel(val int, ch chan int) {
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
