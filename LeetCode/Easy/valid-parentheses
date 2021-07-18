/*
Link: https://leetcode.com/problems/valid-parentheses
91/ 91 test cases passed
Runtime: 0 ms
Memory Usage: 2.2 MB
*/

func isValid(s string) bool {
    matchingString := map[string]string{
        ")":"(",
        "}":"{",
        "]":"[",
    }
    var sliceStr []string
    for _, c := range s {
        switch string(c) {
        case "(" , "{", "[":
            sliceStr = append(sliceStr,string(c))
        case ")" , "}", "]":
            if len(sliceStr) == 0 {
                return false
            }
            if sliceStr[len(sliceStr)-1] != matchingString[string(c)] {
                return false
            }
            sliceStr = sliceStr[:len(sliceStr)-1]
            }
    }
    if len(sliceStr) != 0{
        return false
    }
    return true
}
