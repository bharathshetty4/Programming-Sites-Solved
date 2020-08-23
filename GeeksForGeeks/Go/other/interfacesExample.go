package main

import "fmt"

type Val struct {
	number  int
	strings string
}

type MyInterface interface {
	PrintValues()
}

func (v *Val) PrintValues() {
	fmt.Printf("number: %d, string: %s\n", v.number, v.strings)
}

//func SomeFunction(v Val){
//	//do all Val related operations here
//	v.PrintValues()
//}

func newFunc(v MyInterface){
	v.PrintValues()
}

func main() {
	newVal := Val{10, "ten"}
	//SomeFunction(newVal)
	//you can pass the struct object as an interface,
	//writing UT for interface will be simpler
	newFunc(&newVal)

	val2 := Val{20, "twenty"}
	newFunc(&val2)
}
