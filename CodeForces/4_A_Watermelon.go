package main

/*
URL: https://codeforces.com/problemset/problem/4/A
Status: Accepted
Runtime: 62 ms
Memory Usage: 10400 KB
*/
import (
	"fmt"
)

func main() {
	var kg int
	fmt.Scanf("%d", &kg)
	if kg != 2 && kg%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
