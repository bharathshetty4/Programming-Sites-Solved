//https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem?h_l=interview&playlist_slugs%5B%5D%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D%5B%5D=dictionaries-hashmaps

//not working
package main

import (
	"fmt"
)

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	var ana int32
	var count int
	strMap := make(map[string]int)

	for i := range s {
		strMap[string(s[i])] = strMap[string(s[i])] + 1
		if strMap[string(s[i])] == 2 {
			strMap[string(s[i])] = 0
			count++
		}
	}

	return ana
}

func main() {

	s := "abba"
	result := sherlockAndAnagrams(s)

	fmt.Printf("%d\n", result)

}
