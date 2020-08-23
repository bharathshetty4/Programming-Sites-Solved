//https://www.geeksforgeeks.org/linear-search-using-multi-threading/
package main

import (
	"fmt"
)

func linearSearch(element, key, pos int, chanEle chan int) chan int {
	fmt.Println("pos",pos)
	if element == key {
		chanEle <- pos
	}
	return chanEle
}

func main() {

	chanElem := make(chan int)
	elements := []int{1, 4, 13, 24, 55, 68}

	for pos, elem := range elements {
		go linearSearch(elem, 13, pos, chanElem)
	}
	fmt.Println("position is: ", <-chanElem)

}
