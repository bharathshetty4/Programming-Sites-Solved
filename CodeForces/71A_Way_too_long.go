/*
URL: https://codeforces.com/problemset/problem/71/A
Status: Accepted
Runtime: 62 ms
Memory Usage: 10400 KB
*/
package main

import (
	"fmt"
)

func main() {
	var wLen int
	fmt.Scanf("%d\n", &wLen) // \n is must for codeforces :)
	var wArray []string
	word := ""
	for i := 0; i < wLen; i++ {
		fmt.Scanf("%s\n", &word)
		wArray = append(wArray, word)
	}
	for i := range wArray {
		wLength := len(wArray[i])
		if wLength <= 10 {
			fmt.Println(wArray[i])
			continue
		}
		fmt.Printf("%s%d%s\n", string(wArray[i][0]), wLength-2, string(wArray[i][wLength-1]))
	}
}
