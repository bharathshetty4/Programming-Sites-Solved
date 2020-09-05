package main

/*
URL: https://leetcode.com/problems/find-and-replace-pattern/
Status: Passed
Runtime: 0 ms, faster than 100.00% of Go online submissions for Find and Replace Pattern.
Memory Usage: 2.5 MB, less than 22.58% of Go online submissions for Find and Replace Pattern.
*/
import (
	"fmt"
)

func findAndReplacePattern(words []string, pattern string) []string {
	matchedWords := []string{}
	patternLen := len(pattern)
	//check pattern len also
	if patternLen < 1 || patternLen > 50 {
		return matchedWords
	}

	//check pattern
	letterPositionMap := make(map[byte][]int)
	for i := 0; i < patternLen; i++ {
		letterPositionMap[pattern[i]] = append(letterPositionMap[pattern[i]], i)
	}

	//check each words
	for _, word := range words {
		//check each words
		if len(word) != patternLen || len(word) < 1 || len(word) > 50 {
			continue
		}
		sameWord := true
		wordMap := make(map[byte][]int)
		for i := 0; i < patternLen; i++ {
			wordMap[word[i]] = append(wordMap[word[i]], i)
		}
		for _, pVal := range letterPositionMap {
			if !sameWord {
				break
			}
			for _, wVal := range wordMap {
				if pVal[0] == wVal[0] {
					if len(pVal) != len(wVal) {
						sameWord = false
						break
					}
					for key := range pVal {
						if pVal[key] != wVal[key] {
							sameWord = false
						}
					}
				}
			}
		}
		if sameWord {
			matchedWords = append(matchedWords, word)
		}
	}
	return matchedWords
}

func main() {
	fmt.Println("output: ", findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
}
