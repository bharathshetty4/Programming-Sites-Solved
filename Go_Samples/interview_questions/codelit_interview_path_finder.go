package main

// PROBLEM: For the goven matrix data, we need to check whether the assasin can reach the right corner of the matrix without being
// seen by the guards. X is obstacle and '>' guard can see all the right side cells of the matrix and same work for other 3 types.

import (
	"fmt"
)

func main() {
	fmt.Println("input: <>, output:", Solution([]string{"...Xv", "AX..^", ".XX.."}))
	fmt.Println("input: <>, output:", Solution([]string{"X.....>", "..v..X.", ".>..X..", "A......"}))
}

func Solution(B []string) bool {
	// get the matrix from input
	matrix := [][]string{}
	rowLen := 0
	aRow, aCol := 0, 0
	for _, row := range B {
		var col []string
		for idx, rowChar := range row {
			if string(rowChar) == "A" {
				aRow, aCol = rowLen, idx
			}
			col = append(col, string(rowChar))
		}
		rowLen++
		matrix = append(matrix, col)
	}
	colLen := len(matrix[0])
	// fmt.Println("before mark", matrix, rowLen, colLen, aRow, aCol)

	// update guard range
	var found bool
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			// fmt.Println("matrix[i][j]", matrix[i][j])
			switch matrix[i][j] {
			case "^", "v", ">", "<":
				// fmt.Println("Call markMatrix for ", matrix[i][j])
				matrix, found = markMatrix(matrix, i, j, rowLen, colLen)
				if found {
					return false
				}
			}
		}
	}
	fmt.Println("updated", matrix)

	return pathFinder(matrix, aRow, aCol, rowLen, colLen)
}

func markMatrix(matrix [][]string, row, col, rowLen, colLen int) ([][]string, bool) {
	// based on the char mark the box
	char := matrix[row][col]
	switch char {
	case ">":
		// fmt.Println("Coming markMatrix for ", char)
		for i := col + 1; i < colLen; i++ {
			currentChar := matrix[row][i]
			switch currentChar {
			case "A":
				return matrix, true
			case "^", "v", ">", "<":
				return matrix, false
			}
			matrix[row][i] = "g"
		}
	case "<":
		// fmt.Println("Coming markMatrix for ", char)
		for i := col - 1; i >= 0; i-- {
			currentChar := matrix[row][i]
			switch currentChar {
			case "A":
				return matrix, true
			case "^", "v", ">", "<":
				return matrix, false
			}
			matrix[row][i] = "g"
		}
	case "v":
		// fmt.Println("Coming markMatrix for ", char)
		for j := row + 1; j < rowLen; j++ {
			currentChar := matrix[j][col]
			switch currentChar {
			case "A":
				return matrix, true
			case "^", "v", ">", "<":
				return matrix, false
			}
			matrix[j][col] = "g"
		}
	case "^":
		// fmt.Println("Coming markMatrix for ", char)
		for j := row - 1; j >= 0; j-- {
			currentChar := matrix[j][col]
			switch currentChar {
			case "A":
				return matrix, true
			case "^", "v", ">", "<":
				return matrix, false
			}
			matrix[j][col] = "g"
		}
	}
	return matrix, false
}

func pathFinder(matrix [][]string, startRow, startCol, rowLen, colLen int) bool {
	return findPath(&matrix, rowLen, colLen, startRow, startCol, true)
}

func findPath(matrix *[][]string, rowLen, colLen, i, j int, start bool) bool {
	if i < 0 || j < 0 || i >= rowLen || j >= colLen {
		return false
	}
	currChar := (*matrix)[i][j]
	if currChar == "g" || currChar == "X" || currChar == ">" || currChar == "<" ||
		currChar == "v" || currChar == "^" || (currChar == "A" && !start) {
		return false
	}
	if (*matrix)[i][j] == "." {
		(*matrix)[i][j] = "g"
	}
	if i == rowLen-1 && j == colLen-1 {
		return true
	}
	return findPath(matrix, rowLen, colLen, i, j+1, false) || findPath(matrix, rowLen, colLen, i, j-1, false) ||
		findPath(matrix, rowLen, colLen, i+1, j, false) || findPath(matrix, rowLen, colLen, i-1, j, false)
}
