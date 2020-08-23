package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
)

func MakeRequest(url string, ch chan<-string) {
	start := time.Now()
	resp, _ := http.Get(url)

	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := []string{"https://google.com","https://quora.com"}
	for _,url := range urls{
		go MakeRequest(url, ch)
	}

	for range urls{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}