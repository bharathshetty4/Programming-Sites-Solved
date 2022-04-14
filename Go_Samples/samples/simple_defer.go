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
