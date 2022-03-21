package main

import (
	"fmt"
)

// PROBLEM: For the given string delete the pair of characters if found together.
// i.e in ACDB, CD is a pair. So removing CD we will get AB. Here AB is again a pair which can be removed too, So the
// output should print "".

// input: AB output: ""
// input: ABC output: "C"
// input: CDD output: "D"
// input: ABABCDB output: "B"

// TODO: Use stacks insteda of string check. This solution gave timeout error.
func Solution(S string) string {
	// write your code in Go 1.4
	wordMap := map[string]string{"A": "B", "C": "D"}
	for {
		idx := 0
	label:
		strLen := len(S)
		for ; idx < strLen; idx++ {

			if _, ok := wordMap[string(S[idx])]; ok {
				if idx-1 >= 0 && string(S[idx-1]) == wordMap[string(S[idx])] {
					tempStr := S[0 : idx-1]
					if idx < len(S) {
						S = tempStr + S[idx+1:]
					} else {
						S = tempStr
					}
					idx = idx - 2
					goto label
				}
				if idx+1 < len(S) && string(S[idx+1]) == wordMap[string(S[idx])] {
					tempStr := S[0:idx]
					if idx+1 < len(S) {
						S = tempStr + S[idx+2:]
					} else {
						S = tempStr
					}
					idx = idx - 1
					goto label
				}
			}
		}
		break
	}
	return S
}

func main() {
	fmt.Println("input: <>, output:", Solution("ACD"))
}
