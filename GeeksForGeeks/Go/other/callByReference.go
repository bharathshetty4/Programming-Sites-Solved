// It is not possible to use the same address location in golang,
// So Call by reference doesnot work
package main

import "fmt"

func fn(m *int) {
	fmt.Printf("%v\n",m)
	*m = 1
}

func main() {
	var m int
	fn(&m)
	fmt.Printf("%v\n",&m)
	fmt.Println(m)
}

//output is:
//0xc000006028
//0xc000006028
//1
