package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
)

func main(){

	fmt.Println("Hello!!")
	response, err := http.Get("http://golang.org/")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Println("Response body is ",response.Body)

	defer response.Body.Close()

	byteBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("Response body is", string(byteBody))
}

