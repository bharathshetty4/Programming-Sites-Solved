package main

/*
URL: https://leetcode.com/problems/longest-substring-without-repeating-characters
Status: Success. 987 / 987 test cases passed.
Runtime: 4 ms, faster than 90.16% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 3.2 MB, less than 42.88% of Go online submissions for Longest Substring Without Repeating Characters.
*/
import (
	"fmt"
)

// check each characters and if a character is repeated, set the start index to 'current repeated index + 1' so that
// a new sub-string count can be formed. Calculate longest substring count whenever you see the repeated character.
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	longestSubStr := 1
	startSubStrIndex := 0
	endSubStrIndex := 0
	charIndex := make(map[int32]int)
	for idx, char := range s {
		// a duplicate character found.
		if charOldIndex, ok := charIndex[char]; ok {
			if charOldIndex < startSubStrIndex {
				// if the duplicate char found exist before the start index, nothing to do and continue the computation
			} else {
				currentSubStrLen := endSubStrIndex - startSubStrIndex
				if currentSubStrLen > longestSubStr {
					longestSubStr = currentSubStrLen
				}
				startSubStrIndex = charOldIndex + 1
			}
		}
		// keep incrementing the end index of the substring
		endSubStrIndex++
		charIndex[char] = idx
	}
	currentSubStrLen := endSubStrIndex - startSubStrIndex
	if currentSubStrLen > longestSubStr {
		longestSubStr = currentSubStrLen
	}
	return longestSubStr
}

func main() {
	fmt.Println("input: zabcabcbb, output:", lengthOfLongestSubstring("zabcabcbb"))
	fmt.Println("input: abcabcbb, output:", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("input: ab, output:", lengthOfLongestSubstring("ab"))
	fmt.Println("input: abbac, output:", lengthOfLongestSubstring("abbac"))
}
