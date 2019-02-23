package main

//not working
import (
	"fmt"
)

func minimumSwaps(arr []int32) int32 {
	fmt.Println(arr)
	if len(arr) <= 1 {
		return 0
	}
	var swap int32

	for i:=0;i<len(arr)-1;i++{
		if arr[i] > arr[i+1] {
			for j:=i+1;j<len(arr)-1;j++{
				if arr[j] < arr [j-1] {
					temp := arr[j]
					arr[j] = arr[i]
					arr[i] = temp
					swap++
					i--
					break
				}else {
					for j:=i+1;j<len(arr)-1;j++{
						temp := arr[j]
						arr[j] = arr[j+1]
						arr[j+1] = temp
					}
					swap++
				}
			}
		}
	}

	return swap

}

func main(){
	arr := []int32{2,3,4,1,5}
	res := minimumSwaps(arr)
	fmt.Println(res)

}
