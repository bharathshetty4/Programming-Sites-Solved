package main

import (
  "fmt"
  "sync"
)

func main() {
  var wg = new(sync.WaitGroup)
  wg.Add(1)
  go printHello(wg)
  wg.Wait()
}

func printHello(wg *sync.WaitGroup) {
  for i := 0; i < 5; i++ {
    fmt.Println("Hello, World!")
  }
  wg.Done()
  
}
