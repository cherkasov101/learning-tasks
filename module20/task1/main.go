package main

import "fmt"

const size = 3

func main() {
	var matrix [size][size]int
	fmt.Println("Введите матрицу:")
	for i := 0; i < size*size; i++ {
		row := i / size
		col := i % size
		_, err := fmt.Scan(&matrix[row][col])
		if err != nil {
			fmt.Println("Некорректный ввод")
		}
	}

	fmt.Println(determinant(matrix))
}

func determinant(matrix [size][size]int) (det int) {
	det = matrix[0][0] * matrix[1][1] * matrix[2][2]
	det -= matrix[0][0] * matrix[1][2] * matrix[2][1]
	det -= matrix[0][1] * matrix[1][0] * matrix[2][2]
	det += matrix[0][1] * matrix[1][2] * matrix[2][0]
	det += matrix[0][2] * matrix[1][0] * matrix[2][1]
	det -= matrix[0][2] * matrix[1][1] * matrix[2][0]

	return
}
