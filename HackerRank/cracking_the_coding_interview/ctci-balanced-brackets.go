// https://www.hackerrank.com/challenges/ctci-balanced-brackets/problem
package main

import "fmt"

func main() {
	var count int
	fmt.Scanf("%d", &count)

	var bracketsMap = make(map[string]string)
	bracketsMap["}"] = "{"
	bracketsMap[")"] = "("
	bracketsMap["]"] = "["

	for i := 0; i < count; i++ {
		var tempString string
		var isValid = true
		var bracketsOpen []string

		fmt.Scanf("%s", &tempString)
		if (len(tempString) % 2) != 0 {
			isValid = false
		} else {
			for j := 0; j < len(tempString); j++ {
				if bracketsMap[string(tempString[j])] == "" {
					bracketsOpen = append(bracketsOpen, string(tempString[j]))
				} else {
					if bracketsMap[string(tempString[j])] != bracketsOpen[len(bracketsOpen)-1] {
						isValid = false
						break
					}
					popIndex := len(bracketsOpen) - 1
					bracketsOpen = append(bracketsOpen[:popIndex], bracketsOpen[popIndex+1:]...)

				}

			}
		}
		if isValid {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
