package main

import (
	"errors"
	"fmt"
)

func deferExample() (err error) {
	defer fmt.Println("err1:", err) // nil , err will not get overwritten
	err = errors.New("new err")
	defer func() {
		fmt.Println("err2:", err) // err2: new global err, will get overwritten
	}()
	defer func(err error) {
		err = errors.New("local err")
		fmt.Println("err3:", err) // err3: local err
	}(err)
	err = errors.New("new global err")
	return
}

func main() {
	fmt.Println("Hello", deferExample()) // prints 'Hello new global err'
}

/*
output:
err3: local err
err2: new global err
err1: <nil>
Hello new global err
*/

// only the defer body will get run as part of defer. This doesnot work if we are passing a variable though as updating varaible updates
// value inside defer aswell.
func main() {
	defer func(tt time.Time) {
		fmt.Println("Hello defer", tt) // Hello 2023-03-22 15:22:56.278517 +0100 CET m=+0.000082018
	}(time.Now())
	fmt.Println("Hello", time.Now()) // Hello 2023-03-22 15:22:56.278517 +0100 CET m=+0.000082118
	time.Sleep(1 * time.Second)
	fmt.Println("Hello", time.Now())  // Hello 2023-03-22 15:22:57.279791 +0100 CET m=+1.001350562
}
