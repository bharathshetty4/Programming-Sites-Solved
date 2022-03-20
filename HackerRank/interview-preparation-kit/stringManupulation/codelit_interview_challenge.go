package main

import (
	"fmt"
)

/*
URL:
Status:
Runtime: <> ms, faster than <> of Go online submissions for Add Two Numbers.
Memory Usage: <> MB, less than <> of Go online submissions for Add Two Numbers.
*/

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
