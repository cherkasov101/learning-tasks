package main

import "fmt"

func B(A func(int, int) int) {
	defer fmt.Println("B вызывает A с аргументами 5 и 4, получает:", A(5, 4))
	fmt.Println("Начало работы B")
}

func main() {
	B(func(i int, j int) int {
		i++
		j--
		return i + j
	})

	B(func(i int, j int) int {
		return i*i + j*j
	})

	B(func(i int, j int) int {
		return (i - j) * (i + j)
	})
}
