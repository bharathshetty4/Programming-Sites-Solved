// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var anyVar any
	anyVar = "world"
	fmt.Println("Hello", anyVar.(string))
  //fmt.Println("Hello", anyVar.(int)) //panic: interface conversion: interface {} is string, not int
}
