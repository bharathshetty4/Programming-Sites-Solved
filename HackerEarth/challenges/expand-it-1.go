//https://www.hackerearth.com/challenges/hiring/nutanix-backend-hiring-challenge/algorithm/expand-it-1/
package main

import (
	"fmt"
)

func getCharCount(countString string) int {
	//fmt.Println("Countstring is: ", countString)
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
			charCount[lastChar] = charCount[lastChar] + cnt

			lastChar = string(compStr[i])
			lastIdx = i
		}
		if i == len(compStr)-1 {
			cnt = getCharCount(compStr[lastIdx+1:])
			charCount[lastChar] = charCount[lastChar] + cnt

		}
	}
	return charCount
}

func getChar(charCount [27]int, index []int) {
	tmpA := int('a')
	for _, idx := range index {
		//fmt.Printf("%d:",idx)
		flag := false
		for i := 1; i <= 26; i++ {
			if idx <= 0 {
				fmt.Println(string(tmpA + i - 2))
				flag = true
				break
			}
			idx = idx - charCount[i]
		}
		if !flag {
			if idx <= 0 {
				fmt.Println("z")
			}else {
				fmt.Println("-1")
			}
		}
	}
}

func main() {
	charCount := expandString("a1b5n7b3a8z3")
	var charIdx [27]int
	for i := 1; i <= 26; i++ {
		tmp := int('a')
		charIdx[i] = charCount[string(tmp+i-1)]
	}

	//fmt.Printf("a count is: %v\n", charCount)
	indexes := []int{2,3,4,40,800,21}
	getChar(charIdx, indexes)
}
