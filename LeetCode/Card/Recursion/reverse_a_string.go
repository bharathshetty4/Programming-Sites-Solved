package main

/*
https://leetcode.com/explore/featured/card/recursion-i/250/principle-of-recursion/1440/
Status: Failing
*/
import (
	"fmt"
)

func reverseString(s []byte) {

	reverse(len(s)-1, s)
}

func reverse(strlen int, s []byte) {
	if len(s) == 0 {
		return
	}

	fmt.Printf("idx %d, str %s, full %s\n", len(s)-1, string(s[len(s)-1]), string(s))
	s[0] = s[len(s)-1]
	reverse(strlen, s[:len(s)-1])
}

func main() {
	reverseString([]byte("hello"))
}
