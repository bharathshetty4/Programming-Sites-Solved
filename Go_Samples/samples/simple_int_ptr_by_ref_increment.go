package main

import (
   "fmt"
)

func main() {
	a := 1
	incrInt(&a)
	incrInt(&a)
	fmt.Println("a= ", a)
}

func incrInt(d *int) {
	*d++
}
