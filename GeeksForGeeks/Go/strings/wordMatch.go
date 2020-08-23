package main

import (
	"fmt"
)

func wordMatch(string1, string2 string) []string {
	var matchedString []string
	str2len := len(string2)
	str1len := len(string1)

	for i := 0; i < str1len; i++ {
		if string1[i] == string2[0] {
			if i+str2len > str1len {
				return matchedString
			}
			for j := 1; j < str2len && string1[i+j] == string2[j]; j++ {
				if j == str2len-1 {
					//all string2 are matched now, print till next word
					temp := string1[i : i+str2len]
					i = i + str2len
					for ; i < str1len; i++ {
						if []byte(" ")[0] == string1[i] || []byte(".")[0] == string1[i] {
							break
						}
						temp = temp + string(string1[i])
					}
					matchedString = append(matchedString, temp)
				}
			}
		}
	}
	return matchedString
}

func main() {
	str1 := "this is the best way to get the best knowledge out there."
	str2 := "th"
	fmt.Println(wordMatch(str1, str2))
	str2 = "kno"
	fmt.Println(wordMatch(str1, str2))
	fmt.Print("Closing the main function")
}
