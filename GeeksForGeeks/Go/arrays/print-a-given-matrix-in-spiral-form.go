//https://www.geeksforgeeks.org/print-a-given-matrix-in-spiral-form/
package main

import (
	"fmt"
)

func printInSpiral(matrix [4][4]int, rowLen, columnLen int) {

	for i := 0; i < rowLen/2; i++ {
		//print LR top row
		for lr := i; lr < columnLen-i-1; lr++ {
			fmt.Printf("%d->", matrix[i][lr])
		}

		//print TB right column
		for tb := i; tb < rowLen-i-1; tb++ {
			fmt.Printf("%d->", matrix[tb][columnLen-i-1])
		}

		//print RL below row
		for rl := columnLen - 1 - i; rl > i; rl-- {
			fmt.Printf("%d->", matrix[rowLen-1-i][rl])
		}

		//print BT left column
		for bt := rowLen - 1 - i; bt > i; bt-- {
			fmt.Printf("%d->", matrix[bt][i])
		}
	}

}

func main() {
	var  matrix [4][4]int

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			matrix[i][j] = i*1 + j*2
		}
	}
	fmt.Println(matrix)
	printInSpiral(matrix, 4, 4)
	fmt.Println()

	//matrix2 := make([][]int, 6*6)
	//
	//for i := 0; i < 6; i++ {
	//	for j := 0; j < 6; j++ {
	//		matrix2[i][j] = i*2 + j + 3
	//	}
	//}
	//fmt.Println(matrix2)
	//printInSpiral(matrix2, 6, 6)
	//fmt.Println()

}
