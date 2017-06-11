package main
import "fmt"
//code link: https://www.hackerrank.com/challenges/2d-array
func main() {
    arr := [6][6]int{}
    var maxval int
    for i:=0;i<6;i++{
        for j:=0;j<6;j++{
            fmt.Scanf("%d",&arr[i][j])
        }
    }
    maxval = -10000
    for i:=0;i<4;i++{
        for j:=0;j<4;j++{
            total := (arr[i][j])+(arr[i][j+1])+(arr[i][j+2])+(arr[i+1][j+1])+(arr[i+2][j])+(arr[i+2][j+1])+(arr[i+2][j+2])
            if (total > maxval){
                maxval = total
            }
        }
    }
    fmt.Println(maxval)
}
