package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
)

// Blogs to read:
// https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html

type chanStruct struct {
	intVal int
	strVal string
}

func main() {
	ch := make(chan chanStruct)
	go updateChannel(ch)
	recvdVal := <-ch
	fmt.Println("Hello", recvdVal)
}

func updateChannel(ch chan chanStruct) {
	obj := chanStruct{intVal: 4, strVal: "bks"}
	ch <- obj
}

// another Example
func MakeRequest(url string, ch chan<-string) {
	start := time.Now()
	resp, _ := http.Get(url)

	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}

func main1() {
	start := time.Now()
	ch := make(chan string)
	urls := []string{"https://google.com","https://quora.com"}
	for _,url := range urls{
	    go MakeRequest(url, ch)
	}
	// read for the same number of times, <-ch will hold the line until the data is received
	for range urls{
	    fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
