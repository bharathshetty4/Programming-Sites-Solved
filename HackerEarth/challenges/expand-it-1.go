package main

import (
	"fmt"
)

func getCharCount(countString string) int {
	fmt.Println("Countstring is: ", countString)
	count := 0
	dec := 1
	for i := len(countString) - 1; i >= 0; i-- {
		count = count + int(countString[i]-48)*dec
		dec = dec * 10
	}
	return count
}

func expandString(compStr string) map[string]int {
	var charCount = make(map[string]int)

	lastIdx := 0
	lastChar := string(compStr[0])
	for i := 1; i < len(compStr); i++ {
		cnt := 0
		if string(compStr[i]) >= "a" && string(compStr[i]) <= "z" {
			cnt = getCharCount(compStr[lastIdx+1 : i])
			fmt.Printf("count for %s = %d\n", lastChar, cnt)

			if _, ok := charCount[lastChar]; !ok {
				charCount[lastChar] = cnt
			} else {
				charCount[lastChar] = charCount[lastChar] + cnt
			}

			lastChar = string(compStr[i])
			lastIdx = i
		}
		if i == len(compStr)-1 {
			cnt = getCharCount(compStr[lastIdx+1:])
			fmt.Printf("count for %s = %d\n", lastChar, cnt)
			if _, ok := charCount[lastChar]; !ok {
				charCount[lastChar] = cnt
			} else {
				charCount[lastChar] = charCount[lastChar] + cnt
			}
		}

	}

	return charCount
}

func main() {
	charCount := expandString("a222b56c8a1")
	fmt.Printf("a count is: %v", charCount)
	 
}
