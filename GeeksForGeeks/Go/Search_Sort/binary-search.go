//https://www.geeksforgeeks.org/binary-search/
package main

import (
	"fmt"
)


func binarySearch(elements []int,low, high, key int) int {
	if high < low {
		return -1
	}

	mid := (low +high)/2
	if elements[mid] == key{
		return mid
	}

	if key > elements[mid] {
		return binarySearch(elements,mid+1,high,key)
	}
	return binarySearch(elements,low,mid-1,key)
}

func main()  {
	elems := []int{1,2,3,4,5,6,7,10,20}
	key:= 20
	pos := binarySearch(elems,0,len(elems),key)
	if pos == -1 {
		fmt.Printf("Couldnt find the key %d in element list %+v\n",key,elems)
	}
	fmt.Printf("found the key %d in element list %+v, at position %d\n",key,elems,pos)


	elems = []int{10,30,50,80,80}
	key = 20
	pos = binarySearch(elems,0,len(elems),key)
	if pos == -1 {
		fmt.Printf("Couldnt find the key %d in element list %+v\n",key,elems)
	}
	fmt.Printf("found the key %d in element list %+v, at position %d\n",key,elems,pos)
}