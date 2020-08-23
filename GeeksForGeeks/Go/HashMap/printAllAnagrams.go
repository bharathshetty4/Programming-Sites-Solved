//https://www.geeksforgeeks.org/given-a-sequence-of-words-print-all-anagrams-together/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortWord(str string) string{
	chars := strings.Split(str,"")
	sort.Strings(chars)
	return strings.Join(chars,"")
}

func printAllAnagrams(words []string) {
	wordMap := make(map[string][]int)
	countMap := make(map[int]string)
	for i, word := range words {
		sortedWord := sortWord(word)
		wordMap[sortedWord] =append(wordMap[sortedWord],i)
	}
	maxCount := 0
	for key := range wordMap{
		countMap[len(wordMap[key])] = key
		if maxCount < len(wordMap[key]){
			maxCount = len(wordMap[key])
		}
	}
	for i:= maxCount;i>=1;i--{
		key := countMap[i]
		wordIdxs := wordMap[key]
		for _,idx := range wordIdxs{
			fmt.Print(words[idx]," ")
		}
	}
}

func main() {
	strngs := []string{"cat", "tac", "dog", "god", "act"}
	printAllAnagrams(strngs)
}
