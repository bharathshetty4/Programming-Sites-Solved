//https://www.hackerearth.com/problem/algorithm/xorray
package main

import (
	"fmt"
)

//donot call this function everytime, calculate at once for the maximum and use it every time
func getXOR(digit int) int {
	var xor = make([]int,digit+1)
	xor[1] = 1

	for i := 2; i <= digit; i++ {
		var val = 1
		for j:= 2;j<i;j++{
			val = val ^ xor[j]
		}
		xor[i] = i ^ val
	}

	return xor[digit]
}

//this will take only 0.7 seconds whereas 2 for loop takes 4 seconds
func getXORNotByFor() []int {
	var a = make([]int,1000009)
	var b = make([]int,1000009)

	a[1]=1;
	b[1]=1;
	for i:=2;i<1000009;i++ {
		a[i]=b[i-1]^i;
		b[i]=b[i-1]^a[i];
	}
	return a
}


//this is hackerearth main function
//func main(){
//	var n,val int
//	fmt.Scanf("%d",&n)
//	for i:=0 ; i < n ;i++ {
//		fmt.Scanf("%d",&val)
//		fmt.Println(getXOR(val))
//	}
//}

//this is the modified main function
func main() {
	fmt.Println(getXOR(4))
}
