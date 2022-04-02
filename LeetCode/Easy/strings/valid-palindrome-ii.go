package main

/*
URL: https://leetcode.com/problems/valid-palindrome-ii
Status: Accepted
Runtime: 12 ms, faster than 98.07% of Go online submissions for Valid Palindrome II.
Memory Usage: 6.5 MB, less than 89.07% of Go online submissions for Valid Palindrome II.
*/
func validPalindrome(s string) bool {
    endLen := len(s)-1
    startLen := 0
    var delChar bool
    for {
        if startLen >=  endLen{
            return true
        }
        
        if s[startLen]==s[endLen]{
            startLen++
            endLen--
        }else{
            if delChar{
                return false
            }
            if s[startLen] == s[endLen-1] && s[startLen+1] == s[endLen]{
                // make sure that the remaining string is palindorm, If yes, return true
                // increase the endLen index by 1 as golang will not include the last char if it is not included
                return isPalindrome(s[startLen:endLen]) || isPalindrome(s[startLen+1:endLen+1])
            }else if s[startLen] == s[endLen-1]{
                endLen--
                delChar=true
            }else if s[startLen+1] == s[endLen]{
                startLen++
                delChar=true
            }else{
                return false
            }
            
        }
        
    }
    
    return true
}

func isPalindrome(t string) bool {
    strLen := len(t)
	for i := 0; i < strLen/2; i++ {
		if t[i] != t[strLen-1-i] {
			return false
		}
	}
	return true
}
