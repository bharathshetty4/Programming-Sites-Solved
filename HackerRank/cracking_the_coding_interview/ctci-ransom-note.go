// https://www.hackerrank.com/challenges/ctci-ransom-note/problem
package main
import "fmt"

func main() {
	var magCount, ranCount int
	var tempString string
	var possible = true
	var lettersDict = make(map[string]int)
	fmt.Scanf("%d%d",&magCount,&ranCount)
	for i := 0 ; i < magCount;i++{
		fmt.Scanf("%s",&tempString)
		lettersDict[tempString] = lettersDict[tempString] + 1
	}

	for i := 0 ; i < ranCount;i++{
		fmt.Scanf("%s",&tempString)
		if lettersDict[tempString] <= 0 {
			possible = false
			break
		}else {
			lettersDict[tempString] = lettersDict[tempString] - 1
		}
	}
	if possible {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}