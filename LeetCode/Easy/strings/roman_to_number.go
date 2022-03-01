package main

/*
URL: https://leetcode.com/problems/roman-to-integer
Status: Success
Runtime: 12 ms, faster than 50.68% of Go online submissions for Roman to Integer.
Memory Usage: 3 MB, less than 88.14% of Go online submissions for Roman to Integer.
*/
import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	strLen := len(s)
	sum := 0
	negRoman := map[string]string{"I": "VX", "X": "LC", "C": "DM"}
	romanMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	for idx, strChar := range s {
		// we need to subtract the sum if the next char is in the negative roman map, i.e follows below condition
		/*
		   I can be placed before V (5) and X (10) to make 4 and 9.
		   X can be placed before L (50) and C (100) to make 40 and 90.
		   C can be placed before D (500) and M (1000) to make 400 and 900.
		*/

		if _, ok := negRoman[string(strChar)]; ok && idx+1 < strLen &&
			strings.Contains(negRoman[string(strChar)], string(s[idx+1])) {
			sum = sum - romanMap[string(strChar)]
			continue
		}
		sum = sum + romanMap[string(strChar)]

	}
	return sum
}

func main() {
	fmt.Println("input: <>, output:", romanToInt("III"))
}
