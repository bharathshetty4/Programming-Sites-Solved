package main

import "fmt"

func main() {
	a := 2
	b := &a
	fmt.Println(a) // 2
	fmt.Println(*b) // 2

	*b = 3
	fmt.Println(a) // 3
	fmt.Println(*b) // 3

	a = 4
	fmt.Println(a) // 4
	fmt.Println(*b) // 4
}

/*
Both a and *b refer to the same variable internally.  Hence the changing the value of one reflects in another. 
Also, * and & can be used together as well. But they will cancel out each other.
Hence these are all equivalent,
    a
    *&a
    b
    *&b
    &*b

Note:  *a is not a valid operation as a is not a pointer

*/
