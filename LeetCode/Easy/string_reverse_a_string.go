package main

import "fmt"

/*
URL: https://leetcode.com/problems/reverse-string/
Status: Success
Runtime: 28 ms, faster than 99.13% of Go online submissions for Reverse String.
Memory Usage: 6.6 MB, less than 97.93% of Go online submissions for Reverse String.
*/


// strings are immutable and you cannot do str[i]
func reverseString(s []byte)  {
    strLen := len(s)
    for i:=0; i<strLen/2;i++{
        s[i],s[strLen-1-i] = s[strLen-1-i], s[i]
    }
}

func main() {
	str := "ABCD"
	runes := []rune(str)
	fmt.Println("Hello", str)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	str = string(runes)
	fmt.Println("Hello1", str)
}
