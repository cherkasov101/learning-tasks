package main

import "fmt"

const (
	firstMatrixRows  = 3
	firstMatrixCols  = 5
	secondMatrixRows = 5
	secondMatrixCols = 4
)

func main() {
	fmt.Println("Введите первую матрицу:")
	var firstMatrix [firstMatrixRows][firstMatrixCols]int
	for i := 0; i < firstMatrixRows*firstMatrixCols; i++ {
		row := i / firstMatrixCols
		col := i % firstMatrixCols
		fmt.Scan(&firstMatrix[row][col])
	}

	fmt.Println("Введите вторую матрицу:")
	var secondMatrix [secondMatrixRows][secondMatrixCols]int
	for i := 0; i < secondMatrixRows*secondMatrixCols; i++ {
		row := i / secondMatrixCols
		col := i % secondMatrixCols
		fmt.Scan(&secondMatrix[row][col])
	}

	printMatrix(multiplication(firstMatrix, secondMatrix))
}

func multiplication(firstMatrix [firstMatrixRows][firstMatrixCols]int,
	secondMatrix [secondMatrixRows][secondMatrixCols]int) (resultMatrix [firstMatrixRows][secondMatrixCols]int) {
	for i := 0; i < firstMatrixRows*secondMatrixCols; i++ {
		row := i / secondMatrixCols
		col := i % secondMatrixCols
		for j := 0; j < firstMatrixCols; j++ {
			resultMatrix[row][col] += firstMatrix[row][j] * secondMatrix[j][col]
		}
	}
	return
}

func printMatrix(matrix [firstMatrixRows][secondMatrixCols]int) {
	fmt.Println("Результат умножения матриц:")
	for _, i := range matrix {
		fmt.Println(i)
	}
}
