package Array
//code link: https://www.hackerrank.com/challenges/arrays-ds
import (
	"fmt"
)
func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var size int
	fmt.Scanf("%d", &size)
	vals := make([]string,size)
	for j:= 0 ; j<size;j++{
		fmt.Scanf("%s", &vals[j])
	}
	for j, _ := range vals {
		fmt.Printf("%s ",vals[size-j-1])
	}
}