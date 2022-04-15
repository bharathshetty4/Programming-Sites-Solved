package main

import "fmt"

/*
  Everything in Go is passed by value, slices too. But a slice value is a header,
  describing a contiguous section of a backing array, and a slice value only contains
  a pointer to the array where the elements are actually stored. The slice value does
  not include its elements (unlike arrays).

  So when you pass a slice to a function, a copy will be made from this header, including
  the pointer, which will point to the same backing array. Modifying the elements of the
  slice implies modifying the elements of the backing array, and so all slices which share
  the same backing array will "observe" the change.

  type SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
  }

*/

func updateSlice(mySlice []string) {
	mySlice[0] = "oneupdated"
	mySlice = append(mySlice, "three") // append did not work
	// mySlice[3] = "oneupdated"  //index out of range panic
	myLocalSlice := mySlice
	myLocalSlice[1] = "twoupdated" // will not update the original slice, works with ptr
}

func updateSliceWithPtr(mySlice *[]string) {
	(*mySlice)[0] = "oneptr"
	(*mySlice) = append((*mySlice), "threePtr") // append worked
}

func main() {
	var mySlice = []string{"one", "two"}
	fmt.Println(mySlice) //[one two]
	updateSlice(mySlice)
	fmt.Println(mySlice) // [oneupdated two]
	updateSliceWithPtr(&mySlice)
	fmt.Println(mySlice) // [oneptr two threePtr]
	myMainLocalSlice := mySlice
	myMainLocalSlice[1] = "twoupdated" // will update the original slice value too
	fmt.Println(mySlice)               // [oneptr twoupdated threePtr]
}

/*
op:

[one two]
[oneupdated two]
[oneptr two threePtr]
[oneptr twoupdated threePtr]

*/
